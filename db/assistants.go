package db

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/RunnersRevival/outrun/config"
	"github.com/RunnersRevival/outrun/config/eventconf"
	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db/dbaccess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/netobj/constnetobjs"
	"github.com/RunnersRevival/outrun/obj"

	bolt "go.etcd.io/bbolt"
)

const (
	SessionIDSchema = "OUTRUN_%s"
)

func NewAccountWithID(uid string, resetcount int64) netobj.Player {
	randChar := func(charset string, length int64) string {
		runes := []rune(charset)
		final := make([]rune, 10)
		for i := range final {
			final[i] = runes[rand.Intn(len(runes))]
		}
		return string(final)
	}

	username := ""
	password := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	migrationPassword := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	userPassword := ""
	key := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	playerState := netobj.DefaultPlayerState()
	characterState := netobj.DefaultCharacterState()
	chaoState := constnetobjs.DefaultChaoState()
	mileageMapState := netobj.DefaultMileageMapState()
	mileageFriends := []netobj.MileageFriend{}
	playerVarious := netobj.DefaultPlayerVarious()
	optionUserResult := netobj.DefaultOptionUserResult()
	rouletteInfo := netobj.DefaultRouletteInfo()
	wheelOptions := netobj.DefaultWheelOptions(playerState.NumRouletteTicket, rouletteInfo.RouletteCountInPeriod, enums.WheelRankNormal, consts.RouletteFreeSpins)
	// TODO: get rid of logic here?
	allowedCharacters := []string{}
	allowedChao := []string{}
	for _, chao := range chaoState {
		if chao.Level < 10 { // not max level
			allowedChao = append(allowedChao, chao.ID)
		}
	}
	for _, character := range characterState {
		if character.Star < 10 { // not max star
			allowedCharacters = append(allowedCharacters, character.ID)
		}
	}
	if config.CFile.Debug {
		mileageMapState.Episode = 15
		// testCharacter := netobj.DefaultCharacter(constobjs.CharacterXMasSonic)
		// characterState = append(characterState, testCharacter)
	}
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(playerState, allowedCharacters, allowedChao)
	personalEvents := []eventconf.ConfiguredEvent{}
	messages := []obj.Message{}
	operatorMessages := []obj.OperatorMessage{}
	loginBonusState := netobj.DefaultLoginBonusState(0)
	eventState := netobj.DefaultEventState()
	battleState := netobj.DefaultBattleState()
	disallowInactivePurge := false
	return netobj.NewPlayer(
		uid,
		username,
		password,
		migrationPassword,
		userPassword,
		key,
		playerState,
		characterState,
		chaoState,
		mileageMapState,
		mileageFriends,
		playerVarious,
		optionUserResult,
		wheelOptions,
		rouletteInfo,
		chaoRouletteGroup,
		personalEvents,
		messages,
		operatorMessages,
		loginBonusState,
		false,
		eventState,
		resetcount,
		battleState,
		disallowInactivePurge,
	)
}

func NewAccount() netobj.Player {
	// create ID
	newID := ""
	for i := range make([]byte, 10) {
		if i == 0 { // if first character
			newID += strconv.Itoa(rand.Intn(9) + 1)
		} else {
			newID += strconv.Itoa(rand.Intn(10))
		}
	}
	return NewAccountWithID(newID, 0)
}

func SavePlayer(player netobj.Player) error {
	j, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = dbaccess.Set(consts.DBBucketPlayers, player.ID, j)
	return err
}

func GetPlayer(uid string) (netobj.Player, error) {
	var player netobj.Player
	playerData, err := dbaccess.Get(consts.DBBucketPlayers, uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	err = json.Unmarshal(playerData, &player)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func DeletePlayer(uid string) error {
	err := dbaccess.Delete(consts.DBBucketPlayers, uid)
	return err
}

func GetPlayerBySessionID(sid string) (netobj.Player, error) {
	sidResult, err := dbaccess.Get(consts.DBBucketSessionIDs, sid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	uid, _ := ParseSIDEntry(sidResult)
	player, err := GetPlayer(uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func AssignSessionID(uid string) (string, error) {
	uidB := []byte(uid)
	hash := md5.Sum(uidB)
	hashStr := fmt.Sprintf("%x", hash)
	sid := fmt.Sprintf(SessionIDSchema, hashStr)
	value := fmt.Sprintf("%s/%s", uid, time.Now().Unix()) // register the time that the session ID was assigned
	valueB := []byte(value)
	err := dbaccess.Set(consts.DBBucketSessionIDs, sid, valueB)
	return sid, err
}

func ParseSIDEntry(sidResult []byte) (string, int64) {
	split := strings.Split(string(sidResult), "/")
	uid := split[0]
	timeAssigned, _ := strconv.Atoi(split[1])
	return uid, int64(timeAssigned)
}

func IsValidSessionTime(sessionTime int64) bool {
	timeNow := time.Now().Unix()
	if sessionTime+consts.DBSessionExpiryTime < timeNow {
		return false
	}
	return true
}

func IsValidSessionID(sid []byte) (bool, error) {
	sidResult, err := dbaccess.Get(consts.DBBucketSessionIDs, string(sid))
	if err != nil {
		return false, err
	}
	_, sessionTime := ParseSIDEntry(sidResult)
	return IsValidSessionTime(sessionTime), err
}

func PurgeSessionID(sid string) error {
	err := dbaccess.Delete(consts.DBBucketSessionIDs, sid)
	return err
}

func PurgeAllExpiredSessionIDs() {
	keysToPurge := [][]byte{}
	each := func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(consts.DBBucketSessionIDs))
		err2 := bucket.ForEach(func(k, v []byte) error { // for each value in the session bucket
			_, sessionTime := ParseSIDEntry(v) // get time the session was created
			if !IsValidSessionTime(sessionTime) {
				keysToPurge = append(keysToPurge, k)
			}
			return nil
		})
		return err2
	}
	dbaccess.ForEachLogic(each) // do the logic above
	for _, key := range keysToPurge {
		PurgeSessionID(string(key))
	}
}

func SaveTransferCredentials(transferCreds netobj.TransferCredentials) error {
	j, err := json.Marshal(transferCreds)
	if err != nil {
		return err
	}
	err = dbaccess.Set(consts.DBBucketTransferCreds, transferCreds.TransferID, j)
	return err
}

func GetTransferCredentials(tid string) (netobj.TransferCredentials, error) {
	var transferCreds netobj.TransferCredentials
	trCredsData, err := dbaccess.Get(consts.DBBucketTransferCreds, tid)
	if err != nil {
		return netobj.DefaultTransferCredentials(), err
	}
	err = json.Unmarshal(trCredsData, &transferCreds)
	if err != nil {
		return netobj.DefaultTransferCredentials(), err
	}
	return transferCreds, nil
}

func DeleteTransferCredentials(tid string) error {
	err := dbaccess.Delete(consts.DBBucketTransferCreds, tid)
	return err
}

func BattleGetPlayer(uid string) (netobj.Player, error) {
	var player netobj.Player
	playerData, err := dbaccess.Get(consts.BattleDBBucketWaiting, uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	err = json.Unmarshal(playerData, &player)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func BattleGetMatchedPoolPlayer(uid string) (netobj.Player, error) {
	var player netobj.Player
	playerData, err := dbaccess.Get(consts.BattleDBBucketMatched, uid)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	err = json.Unmarshal(playerData, &player)
	if err != nil {
		return constnetobjs.BlankPlayer, err
	}
	return player, nil
}

func BattleSaveWaitingPlayer(player netobj.Player) error {
	j, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = dbaccess.BattleDBSet(consts.BattleDBBucketWaiting, player.ID, j)
	return err
}

func BattleSaveMatchedPlayer(player netobj.Player) error {
	j, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = dbaccess.BattleDBSet(consts.BattleDBBucketMatched, player.ID, j)
	return err
}

func BattleDeleteWaitingPlayer(uid string) error {
	err := dbaccess.BattleDBDelete(consts.BattleDBBucketWaiting, uid)
	return err
}

func BattleDeleteMatchedPlayer(uid string) error {
	err := dbaccess.BattleDBDelete(consts.BattleDBBucketMatched, uid)
	return err
}

func BattleDeletePlayer(uid string) error {
	err := dbaccess.BattleDBDelete(consts.BattleDBBucketWaiting, uid)
	err = dbaccess.BattleDBDelete(consts.BattleDBBucketMatched, uid)
	return err
}
