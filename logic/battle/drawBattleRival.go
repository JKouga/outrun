package battle

import (
	"log"
	"math/rand"
	"time"

	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/db/dbaccess"
	"github.com/RunnersRevival/outrun/netobj"
)

// TODO: Add stricter algorithm for the more expensive "Find a Beatable Opponent" reroll option
func DrawBattleRival(player netobj.Player, limit int) netobj.BattleState {
	if !player.BattleState.MatchedUpWithRival { // are we not matched up yet?
		db.BattleSaveWaitingPlayer(player) // Save player in the waiting pool
		playerIDs := []string{}
		dbaccess.BattleDBForEachKey(consts.BattleDBBucketWaiting, func(k, v []byte) error {
			playerIDs = append(playerIDs, string(k))
			return nil
		})
		currentTime := time.Now().UTC().Unix()
		rivalID := ""
		iterations := 0
		for len(playerIDs) > 0 {
			index := rand.Intn(len(playerIDs))
			potentialRival, err := db.GetPlayer(playerIDs[index])
			if err != nil {
				log.Printf("[WARN] (battle.DrawBattleRival) Unable to get player '%s': %s", playerIDs[index], err.Error())
			} else {
					if player.ID != playerIDs[index] &&
						potentialRival.BattleState.ScoreRecordedToday &&
						!potentialRival.BattleState.MatchedUpWithRival &&
						currentTime < potentialRival.BattleState.BattleEndsAt &&
						(potentialRival.BattleState.PrevDailyBattleHighScore < player.BattleState.PrevDailyBattleHighScore - 1500000) && 
						potentialRival.BattleState.PrevDailyBattleHighScore > player.BattleState.PrevDailyBattleHighScore + 1500000{
						rivalID = playerIDs[index]
						break
					} else { 
						if player.ID != playerIDs[index] && // Player is none of the above, default to fully randomized rival
						potentialRival.BattleState.ScoreRecordedToday &&
						!potentialRival.BattleState.MatchedUpWithRival &&
						currentTime < potentialRival.BattleState.BattleEndsAt {
							rivalID = playerIDs[index]
							break
					}
				}
			}
			playerIDs[index] = playerIDs[len(playerIDs)-1]
			playerIDs = playerIDs[:len(playerIDs)-1]
			iterations++
			if iterations >= limit && limit > 0 {
				break
			}
		}
		if len(rivalID) > 0 {
			rival, err := db.GetPlayer(rivalID)
			if err != nil {
				log.Printf("[WARN] (battle.DrawBattleRival) Unable to get player '%s': %s", rivalID, err.Error())
			} else {
				rival.BattleState.RivalID = player.ID
				rival.BattleState.MatchedUpWithRival = true
				db.BattleDeleteWaitingPlayer(player.ID)
				db.BattleDeleteWaitingPlayer(rival.ID)
				db.BattleSaveMatchedPlayer(player)
				db.BattleSaveMatchedPlayer(rival)
				err = db.SavePlayer(rival)
				if err != nil {
					log.Printf("[WARN] (battle.DrawBattleRival) Unable to save rival data: %s", err.Error())
				} else {
					player.BattleState.RivalID = rivalID
					player.BattleState.MatchedUpWithRival = true
					db.BattleDeleteWaitingPlayer(rival.ID)
					db.BattleDeleteWaitingPlayer(rival.ID)
					db.BattleSaveMatchedPlayer(player)
					db.BattleSaveMatchedPlayer(rival)
				}
			}
		}
	} else {
		db.BattleSaveMatchedPlayer(player)
		db.BattleDeleteWaitingPlayer(player.ID) // Player is matched, delete entry from the waiting pool
	}
	if player.ID == player.BattleState.RivalID && player.BattleState.MatchedUpWithRival {
		log.Printf("[WARN] (battle.DrawBattleRival) Somehow matched up with self! Removing match...")
		player.BattleState.MatchedUpWithRival = false
		db.BattleDeleteMatchedPlayer(player.ID)
		db.BattleSaveWaitingPlayer(player)
	}
	return player.BattleState
}
