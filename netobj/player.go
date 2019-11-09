package netobj

import (
	"fmt"
	"time"

	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
)

type Player struct {
	ID                string `json:"userID"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Key               string `json:"key"`
	LastLogin         int64
	PlayerState       PlayerState       `json:"playerState"`
	CharacterState    []Character       `json:"characterState"`
	ChaoState         []Chao            `json:"chaoState"`
	MileageMapState   MileageMapState   `json:"mileageMapState"`
	MileageFriends    []MileageFriend   `json:"mileageFriendList"`
	PlayerVarious     PlayerVarious     `json:"playerVarious"`
	LastWheelOptions  WheelOptions      `json:"ORN_wheelOptions"` // TODO: Make RouletteGroup to hold LastWheelOptions and RouletteInfo?
	RouletteInfo      RouletteInfo      `json:"ORN_rouletteInfo"`
	ChaoRouletteGroup ChaoRouletteGroup `json:"ORN_chaoRouletteGroup"`
}

func NewPlayer(id, username, password, key string, playerState PlayerState, characterState []Character, chaoState []Chao, mileageMapState MileageMapState, mf []MileageFriend, playerVarious PlayerVarious, wheelOptions WheelOptions, rouletteInfo RouletteInfo, chaoRouletteGroup ChaoRouletteGroup) Player {
	return Player{
		id,
		username,
		password,
		key,
		time.Now().Unix(),
		playerState,
		characterState,
		chaoState,
		mileageMapState,
		mf,
		playerVarious,
		wheelOptions,
		rouletteInfo,
		chaoRouletteGroup,
	}
}

/*
func (p *Player) Save() {

}
*/
// TODO: remove any functions that access p.PlayerState since we are not calling from a pointer anyways and it will not modify the object
func (p *Player) AddRings(amount int64) {
	ps := p.PlayerState
	ps.NumRings += amount
}
func (p *Player) SubRings(amount int64) {
	ps := p.PlayerState
	ps.NumRings -= amount
}
func (p *Player) AddRedRings(amount int64) {
	ps := p.PlayerState
	ps.NumRedRings += amount
}
func (p *Player) SubRedRings(amount int64) {
	ps := p.PlayerState
	ps.NumRedRings -= amount
}
func (p *Player) SetUsername(username string) {
	p.Username = username
}
func (p *Player) SetPassword(password string) {
	p.Password = password
}
func (p *Player) AddEnergy(amount int64) {
	ps := p.PlayerState
	ps.Energy += amount
}
func (p *Player) SubEnergy(amount int64) {
	ps := p.PlayerState
	ps.Energy -= amount
}
func (p *Player) SetMainCharacter(cid string) {
	ps := p.PlayerState
	ps.MainCharaID = cid
}
func (p *Player) SetSubCharacter(cid string) {
	ps := p.PlayerState
	ps.SubCharaID = cid
}
func (p *Player) SetMainChao(chid string) {
	ps := p.PlayerState
	ps.MainChaoID = chid
}
func (p *Player) SetSubChao(chid string) {
	ps := p.PlayerState
	ps.SubChaoID = chid
}
func (p *Player) AddItem(item obj.Item) {
	ps := p.PlayerState
	ps.Items = append(ps.Items, item)
}
func (p *Player) RemoveItemOf(iid string) bool {
	newItems := []obj.Item{}
	foundItem := false
	ps := p.PlayerState
	for _, item := range ps.Items {
		if item.ID != iid || foundItem {
			newItems = append(newItems, item)
		} else if !foundItem {
			foundItem = true
		}
	}
	ps.Items = newItems
	return foundItem
}
func (p *Player) IndexOfItem(iid string) int {
	for i, item := range p.PlayerState.Items {
		if item.ID == iid {
			return i
		}
	}
	return -1
}
func (p *Player) RemoveAllItemsOf(iid string) {
	for p.RemoveItemOf(iid) {
	}
}
func (p *Player) AddAnimals(amount int64) {
	ps := p.PlayerState
	ps.Animals += amount
}
func (p *Player) SubAnimals(amount int64) {
	ps := p.PlayerState
	ps.Animals -= amount
}
func (p *Player) ApplyHighScore(score int64) bool {
	ps := p.PlayerState
	if ps.HighScore < score {
		ps.HighScore = score
		return true
	}
	return false
}
func (p *Player) AddDistance(amount int64) {
	ps := p.PlayerState
	ps.TotalDistance += amount
	p.ApplyHighDistance(amount)
}
func (p *Player) ApplyHighDistance(amount int64) {
	ps := p.PlayerState
	ps.HighDistance = amount
}
func (p *Player) AddNewChaoByID(chid string) bool {
	chao := constobjs.Chao[chid]
	netchao := NewNetChao(
		chao,
		enums.ChaoStatusOwned, // TODO: does the idea that a chao is owned mean that it's possible to send chao that are not owned?
		1,
		enums.ChaoDealingNone,
		1, // implies that adding means acquired. This may not be the case if we can send non-owned chao.
	)
	return p.AddNetChao(netchao)
}
func (p *Player) AddNewChao(chao obj.Chao) bool {
	netchao := NewNetChao(
		chao,
		enums.ChaoStatusOwned,
		1,
		enums.ChaoDealingNone,
		1,
	)
	return p.AddNetChao(netchao)
}
func (p *Player) AddNetChao(netchao Chao) bool {
	// Returns whether or not the Chao was already found.
	// It will not add Chao already in the ChaoState.
	if !p.HasChao(netchao.Chao.ID) {
		p.ChaoState = append(p.ChaoState, netchao)
		return false
	}
	return true
}
func (p *Player) HasChao(chid string) bool {
	for _, netchao := range p.ChaoState {
		if netchao.Chao.ID == chid {
			return true
		}
	}
	return false
}
func (p *Player) GetChara(cid string) (Character, error) {
	var char Character
	found := false
	for _, c := range p.CharacterState {
		if c.ID == cid {
			char = c
			found = true
		}
	}
	if !found {
		return char, fmt.Errorf("character not found")
	}
	return char, nil
}
func (p *Player) IndexOfChara(cid string) int {
	for i, char := range p.CharacterState {
		if char.ID == cid {
			return i
		}
	}
	return -1
}
func (p *Player) GetChao(chid string) (Chao, error) {
	var chao Chao
	found := false
	for _, c := range p.ChaoState {
		if c.ID == chid {
			chao = c
			found = true
		}
	}
	if !found {
		return chao, fmt.Errorf("chao not found")
	}
	return chao, nil
}
func (p *Player) IndexOfChao(chid string) int {
	for i, chao := range p.ChaoState {
		if chao.ID == chid {
			return i
		}
	}
	return -1
}
func (p *Player) GetMainChara() (Character, error) {
	ps := p.PlayerState
	cid := ps.MainCharaID
	char, err := p.GetChara(cid)
	return char, err
}
func (p *Player) GetSubChara() (Character, error) {
	ps := p.PlayerState
	cid := ps.SubCharaID
	char, err := p.GetChara(cid)
	return char, err
}
func (p *Player) GetMainChao() (Chao, error) {
	ps := p.PlayerState
	chid := ps.MainChaoID
	chao, err := p.GetChao(chid)
	return chao, err
}
func (p *Player) GetSubChao() (Chao, error) {
	ps := p.PlayerState
	chid := ps.SubChaoID
	chao, err := p.GetChao(chid)
	return chao, err
}
func (p *Player) GetMaxLevelChao() []Chao {
	mxlvl := []Chao{}
	for _, c := range p.ChaoState {
		if c.Level >= 10 { // if max level (or above)
			mxlvl = append(mxlvl, c)
		}
	}
	return mxlvl
}
func (p *Player) GetMaxLevelChaoIDs() []string {
	chao := p.GetMaxLevelChao()
	ids := []string{}
	for _, c := range chao {
		if c.Level >= 10 { // if max level (or above)
			ids = append(ids, c.ID)
		}
	}
	return ids
}
func (p *Player) GetMaxLevelCharacters() []Character {
	mxlvl := []Character{}
	for _, ch := range p.CharacterState {
		if ch.Star >= 10 { // if max stars (or above)
			mxlvl = append(mxlvl, ch)
		}
	}
	return mxlvl
}
func (p *Player) GetMaxLevelCharacterIDs() []string {
	chars := p.GetMaxLevelCharacters()
	ids := []string{}
	for _, ch := range chars {
		//if ch.Level >= 100 { // if max level (or above)
		if ch.Star >= 10 {
			ids = append(ids, ch.ID)
		}
	}
	return ids
}
func (p *Player) GetAllMaxLevelIDs() []string {
	chars := p.GetMaxLevelCharacterIDs()
	chao := p.GetMaxLevelChaoIDs()
	combined := []string{}
	combined = append(combined, chars...)
	combined = append(combined, chao...)
	return combined
}
func (p *Player) AllChaoMaxLevel() bool {
	for _, c := range p.ChaoState {
		if c.Level < 10 {
			return false
		}
	}
	return true
}
func (p *Player) AllCharactersMaxLevel() bool {
	for _, ch := range p.CharacterState {
		if ch.Star < 10 {
			return false
		}
	}
	return true
}
func (p *Player) GetAllNonMaxedChaoAndCharacters() []string {
	result := append(p.GetAllNonMaxedChao(), p.GetAllNonMaxedCharacters()...) // combine two
	return result
}
func (p *Player) GetAllNonMaxedChao() []string {
	result := []string{}
	for _, chao := range p.ChaoState {
		if chao.Level < 10 { // not max level
			result = append(result, chao.ID)
		}
	}
	return result
}
func (p *Player) GetAllNonMaxedCharacters() []string {
	result := []string{}
	for _, character := range p.CharacterState {
		if character.Star < 10 { // not max star
			result = append(result, character.ID)
		}
	}
	return result
}
