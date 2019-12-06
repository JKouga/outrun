package rpcobj

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Mtbcooler/outrun/config/gameconf"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/logic"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/netobj/constnetobjs"
	"github.com/Mtbcooler/outrun/obj/constobjs"
)

func (t *Toolbox) Debug_GetCampaignStatus(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = strconv.Itoa(int(player.MileageMapState.Chapter)) + "," + strconv.Itoa(int(player.MileageMapState.Episode)) + "," + strconv.Itoa(int(player.MileageMapState.Point))
	return nil
}

func (t *Toolbox) Debug_GetAllPlayerIDs(nothing bool, reply *ToolboxReply) error {
	playerIDs := []string{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	final := strings.Join(playerIDs, ",")
	reply.Status = StatusOK
	reply.Info = final
	return nil
}

func (t *Toolbox) Debug_ResetPlayer(uid string, reply *ToolboxReply) error {
	_ = db.NewAccountWithID(uid)
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_GetRouletteInfo(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	rouletteInfo := player.RouletteInfo
	jri, err := json.Marshal(rouletteInfo)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to marshal RouletteInfo: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = string(jri)
	return nil
}

func (t *Toolbox) Debug_ResetChaoRouletteGroup(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(player.PlayerState, player.GetAllNonMaxedCharacters(), player.GetAllNonMaxedChao())
	player.ChaoRouletteGroup = chaoRouletteGroup
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_ResetCharactersAndCompensate(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	toAdd := int64(0)
	for _, char := range player.CharacterState {
		toAdd += char.Level * 7
	}
	player.PlayerState.NumRedRings += toAdd
	player.CharacterState = netobj.DefaultCharacterState()
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_ResetChao(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	player.ChaoState = constnetobjs.NetChaoList
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_MigrateUser(uidToUID string, reply *ToolboxReply) error {
	uidSrc := strings.Split(uidToUID, "->")
	if len(uidSrc) != 2 {
		reply.Status = StatusOtherError
		reply.Info = "improperly formatted string (Example: 1234567890->1987654321)"
	}

	fromUID := uidSrc[0]
	toUID := uidSrc[1]
	oldPlayer, err := db.GetPlayer(fromUID)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	currentPlayer, err := db.GetPlayer(toUID)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	currentPlayer.PlayerState = oldPlayer.PlayerState
	currentPlayer.Username = oldPlayer.Username
	currentPlayer.LastLogin = oldPlayer.LastLogin
	currentPlayer.CharacterState = oldPlayer.CharacterState
	currentPlayer.ChaoState = oldPlayer.ChaoState
	currentPlayer.MileageMapState = oldPlayer.MileageMapState
	currentPlayer.MileageFriends = oldPlayer.MileageFriends
	currentPlayer.PlayerVarious = oldPlayer.PlayerVarious
	currentPlayer.LastWheelOptions = oldPlayer.LastWheelOptions
	currentPlayer.ChaoRouletteGroup = oldPlayer.ChaoRouletteGroup
	currentPlayer.RouletteInfo = oldPlayer.RouletteInfo

	err = db.SavePlayer(currentPlayer)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}

	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_UsernameSearch(username string, reply *ToolboxReply) error {
	playerIDs := []string{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	sameUsernames := []string{}
	for _, uid := range playerIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = "error getting ID " + uid + ": " + err.Error()
			return err
		}
		if player.Username == username {
			sameUsernames = append(sameUsernames, player.ID)
		}
	}
	if len(sameUsernames) == 0 {
		reply.Status = StatusOtherError
		reply.Info = "unable to find ID for username " + username
		return nil
	}
	reply.Status = StatusOK
	reply.Info = strings.Join(sameUsernames, ",")
	return nil
}

func (t *Toolbox) Debug_RawPlayer(uid string, reply *ToolboxReply) error {
	playerSrc, err := dbaccess.Get(consts.DBBucketPlayers, uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = string(playerSrc)
	return nil
}

func (t *Toolbox) Debug_ResetCharacterState(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	player.CharacterState = netobj.DefaultCharacterState()
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_MatchPlayersToGameConf(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		player.CharacterState = netobj.DefaultCharacterState() // already uses AllCharactersUnlocked
		player.ChaoState = constnetobjs.DefaultChaoState()     // already uses AllChaoUnlocked
		player.PlayerState.MainCharaID = gameconf.CFile.DefaultMainCharacter
		player.PlayerState.SubChaoID = gameconf.CFile.DefaultSubChao
		player.PlayerState.MainChaoID = gameconf.CFile.DefaultMainChao
		player.PlayerState.SubCharaID = gameconf.CFile.DefaultSubCharacter
		player.PlayerState.NumRings = gameconf.CFile.StartingRings
		player.PlayerState.NumRedRings = gameconf.CFile.StartingRedRings
		player.PlayerState.Energy = gameconf.CFile.StartingEnergy
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_PrepTag1p0(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	sqrt := func(n int64) int64 {
		fn := float64(n)
		result := math.Sqrt(fn)
		return int64(result)
	}

    for _, uid := range allUIDs {
        player, err := db.GetPlayer(uid)
        if err != nil {
            reply.Status = StatusOtherError
            reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
            return err
        }
        // compensate first
        amountPerLevel := int64(7)  // Red Rings offered per level
        newRedRingAmount := int64(0)
        for _, char := range player.CharacterState {
            newRedRingAmont += char.Level * amountPerLevel
        }
        player.CharacterState = netobj.DefaultCharacterState() // already uses AllCharactersUnlocked
        player.ChaoState = constnetobjs.DefaultChaoState()     // already uses AllChaoUnlocked
        player.PlayerState.MainCharaID = gameconf.CFile.DefaultMainCharacter
        player.PlayerState.SubChaoID = gameconf.CFile.DefaultSubChao
        player.PlayerState.MainChaoID = gameconf.CFile.DefaultMainChao
        player.PlayerState.SubCharaID = gameconf.CFile.DefaultSubCharacter
        player.PlayerState.NumRings = int64(HOWEVER_MANY_RINGS_HERE)
        player.PlayerState.NumRedRings = newRedRingAmount
        player.PlayerState.Energy = gameconf.CFile.StartingEnergy
        player.PlayerState.Items = constobjs.DefaultPlayerStateItems

        err = db.SavePlayer(player)
        if err != nil {
            reply.Status = StatusOtherError
            reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
            return err
        }
    }
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_PlayersByPassword(password string, reply *ToolboxReply) error {
	foundPlayers, err := logic.FindPlayersByPassword(password, false)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "error finding players by password: " + err.Error()
		return err
	}
	playerIDs := []string{}
	for _, player := range foundPlayers {
		playerIDs = append(playerIDs, player.ID)
	}
	final := strings.Join(playerIDs, ",")
	reply.Status = StatusOK
	reply.Info = final
	return nil
}

func (t *Toolbox) Debug_ResetPlayersRank(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		player.PlayerState.Rank = 0 // for some reason, this gets incremented 1 by the game
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_FixWerehogRedRings(uids string, reply *ToolboxReply) error {
	wh := constobjs.CharacterWerehog
	whid := wh.ID
	whrr := wh.PriceRedRings
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		i := player.IndexOfChara(whid)
		if i == -1 {
			reply.Status = StatusOK
			reply.Info = "index not found!"
			return fmt.Errorf("index not found!")
		}
		player.CharacterState[i].Character.PriceRedRings = whrr
		player.CharacterState[i].PriceRedRings = whrr
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}
