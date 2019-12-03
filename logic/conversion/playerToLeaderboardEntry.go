package conversion

import (
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/netobj"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/jinzhu/now"
)

func PlayerToLeaderboardEntry(player netobj.Player, mode int64) obj.LeaderboardEntry {
	friendID := player.ID
	name := player.Username
	url := player.Username + "_findme" // TODO: only used for testing right now
	grade := int64(1)                  // TODO: make this specifiable in the future (this is the high score ranking it seems)
	exposeOnline := int64(0)
	rankingScore := player.PlayerState.HighScore
	rankChanged := int64(0)
	isSentEnergy := int64(0)
	expireTime := now.EndOfWeek().UTC().Unix()
	numRank := player.PlayerState.Rank
	loginTime := player.LastLogin
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := player.CharacterState[player.IndexOfChara(mainCharaID)].Level // TODO: is this right?
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := player.CharacterState[player.IndexOfChara(subCharaID)].Level
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := player.ChaoState[player.IndexOfChao(mainChaoID)].Level
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := player.ChaoState[player.IndexOfChao(subChaoID)].Level
	language := int64(enums.LangEnglish)
	league := player.PlayerState.RankingLeague
	maxScore := player.PlayerState.HighScore
	if mode == 1 { //timed mode?
		rankingScore = player.PlayerState.TimedHighScore
		league = player.PlayerState.QuickRankingLeague
		maxScore = player.PlayerState.TimedHighScore
	}
	return obj.LeaderboardEntry{
		friendID,
		name,
		url,
		grade,
		exposeOnline,
		rankingScore,
		rankChanged,
		isSentEnergy,
		expireTime,
		numRank,
		loginTime,
		mainCharaID,
		mainCharaLevel,
		subCharaID,
		subCharaLevel,
		mainChaoID,
		mainChaoLevel,
		subChaoID,
		subChaoLevel,
		language,
		league,
		maxScore,
	}
}
