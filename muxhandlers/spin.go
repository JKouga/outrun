package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/RunnersRevival/outrun/analytics"
	"github.com/RunnersRevival/outrun/analytics/factors"
	"github.com/RunnersRevival/outrun/config"
	"github.com/RunnersRevival/outrun/config/campaignconf"
	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/db"
	"github.com/RunnersRevival/outrun/emess"
	"github.com/RunnersRevival/outrun/enums"
	"github.com/RunnersRevival/outrun/helper"
	"github.com/RunnersRevival/outrun/logic"
	"github.com/RunnersRevival/outrun/logic/conversion"
	"github.com/RunnersRevival/outrun/logic/roulette"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/RunnersRevival/outrun/requests"
	"github.com/RunnersRevival/outrun/responses"
	"github.com/RunnersRevival/outrun/status"
)

func GetWheelOptions(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	//player.LastWheelOptions = netobj.DefaultWheelOptions(player.PlayerState) // generate new wheel for 'reroll' mechanic
	helper.DebugOut("Time now: %v", time.Now().Unix())
	helper.DebugOut("RoulettePeriodEnd: %v", player.RouletteInfo.RoulettePeriodEnd)
	// check if we need to reset the end period
	endPeriod := player.RouletteInfo.RoulettePeriodEnd
	if time.Now().Unix() > endPeriod {
		player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
	}

	// refresh wheel
	player.LastWheelOptions = logic.WheelRefreshLogic(player, player.LastWheelOptions)

	response := responses.WheelOptions(baseInfo, player.LastWheelOptions)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func CommitWheelSpin(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.CommitWheelSpinRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	helper.DebugOut("request.Count: %v", request.Count)

	freeSpins := consts.RouletteFreeSpins
	campaignList := []obj.Campaign{}
	if campaignconf.CFile.AllowCampaigns {
		for _, confCampaign := range campaignconf.CFile.CurrentCampaigns {
			newCampaign := conversion.ConfiguredCampaignToCampaign(confCampaign)
			campaignList = append(campaignList, newCampaign)
		}
	}
	for index := range campaignList {
		if obj.IsCampaignActive(campaignList[index]) && campaignList[index].Type == enums.CampaignTypeFreeWheelSpinCount {
			freeSpins = campaignList[index].Content
		}
		index++
	}

	endPeriod := player.RouletteInfo.RoulettePeriodEnd
	helper.DebugOut("Time now: %v", time.Now().UTC().Unix())
	helper.DebugOut("End period: %v", endPeriod)
	if time.Now().UTC().Unix() > endPeriod {
		player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
		helper.DebugOut("New roulette period")
		helper.DebugOut("RouletteCountInPeriod: %v", player.RouletteInfo.RouletteCountInPeriod)
	}

	responseStatus := status.OK
	hasTickets := player.PlayerState.NumRouletteTicket > 0
	hasFreeSpins := player.RouletteInfo.RouletteCountInPeriod < freeSpins
	helper.DebugOut("Has tickets: %v", hasTickets)
	helper.DebugOut("Number of tickets: %v", player.PlayerState.NumRouletteTicket)
	helper.DebugOut("Has free spins: %v", hasFreeSpins)
	helper.DebugOut("Roulette count: %v", player.RouletteInfo.RouletteCountInPeriod)
	if hasTickets || hasFreeSpins {
		//if player.LastWheelOptions.NumRemainingRoulette > 0 {
		wonItem := player.LastWheelOptions.Items[player.LastWheelOptions.ItemWon]
		itemExists := player.IndexOfItem(wonItem) != -1
		oldRanking := player.LastWheelOptions.RouletteRank
		if itemExists {
			amountOfItemWon := player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
			helper.DebugOut("wonItem: %v", wonItem)
			helper.DebugOut("amountOfItemWon: %v", amountOfItemWon)
			itemIndex := player.IndexOfItem(wonItem)
			helper.DebugOut("Amount of item player has: %v", player.PlayerState.Items[itemIndex].Amount)
			player.PlayerState.Items[itemIndex].Amount += amountOfItemWon
			helper.DebugOut("New amount of item player has: %v", player.PlayerState.Items[itemIndex].Amount)
		} else {
			if wonItem == strconv.Itoa(enums.IDTypeItemRouletteWin) && oldRanking == enums.WheelRankSuper {
				// Jackpot
				player.PlayerState.NumRings += player.LastWheelOptions.NumJackpotRing
			} else if wonItem == strconv.Itoa(enums.ItemIDRedRing) {
				// Red rings
				player.PlayerState.NumRedRings += player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
			} else if wonItem == strconv.Itoa(enums.ItemIDRing) {
				// Rings
				player.PlayerState.NumRings += player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
			} else if wonItem[:2] == "40" {
				// Chao
				amountOfItemWon := player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
				helper.DebugOut("wonItem: %v", wonItem)
				helper.DebugOut("amountOfItemWon: %v", amountOfItemWon)

				chaoIndex := player.IndexOfChao(wonItem)
				if chaoIndex == -1 { // chao index not found, should never happen
					helper.InternalErr("cannot get index of chao '"+strconv.Itoa(chaoIndex)+"'", err)
					return
				}
				if player.ChaoState[chaoIndex].Status == enums.ChaoStatusNotOwned {
					// earn the Chao
					player.ChaoState[chaoIndex].Status = enums.ChaoStatusOwned
					player.ChaoState[chaoIndex].Acquired = 1
					player.ChaoState[chaoIndex].Level = 0
				}
				//player.ChaoState[chaoIndex].Level += amountOfItemWon
				maxChaoLevel := int64(10)
				if request.Version == "1.1.4" {
					maxChaoLevel = int64(5)
				}
				if player.ChaoState[chaoIndex].Level < maxChaoLevel {
					player.ChaoState[chaoIndex].Level += amountOfItemWon
					if player.ChaoState[chaoIndex].Level > maxChaoLevel { // if max chao level (https://www.deviantart.com/vocaloidbrsfreak97/journal/So-Sonic-Runners-just-recently-updated-574789098)
						excess := player.ChaoState[chaoIndex].Level - maxChaoLevel    // get amount gone over
						amountOfItemWon -= excess                                     // shave it from prize level
						player.ChaoState[chaoIndex].Level = maxChaoLevel              // reset to maximum
						player.ChaoState[chaoIndex].Status = enums.ChaoStatusMaxLevel // set status to MaxLevel
					}
				} else {
					player.ChaoState[chaoIndex].Level = maxChaoLevel              // reset to maximum
					player.ChaoState[chaoIndex].Status = enums.ChaoStatusMaxLevel // set status to MaxLevel
					player.PlayerState.ChaoEggs += 3                              // maxed out; give 3 special eggs as compensation
					
					// refresh the Premium Roulette - should fix issue #19 (https://github.com/RunnersRevival/revival_issues/issues/19)

					chaoCanBeLevelled := !player.AllChaoMaxLevel()
					charactersCanBeLevelled := !player.AllCharactersMaxLevel()
					fixRarities := func(rarities []int64) ([]int64, bool) {
						newRarities := []int64{}
						if !chaoCanBeLevelled && !charactersCanBeLevelled {
							// Wow, they can't upgrade _anything!_
							return newRarities, false
						}
						if config.CFile.Debug {
							player.PlayerState.NumRedRings += 150
							//return []int64{100, 100, 100, 100, 100, 100, 100, 100}, true
							return []int64{0, 0, 0, 0, 0, 0, 0, 0}, true
						}
						for _, r := range rarities {
							if r == 0 || r == 1 || r == 2 { // Chao
								if chaoCanBeLevelled {
									newRarities = append(newRarities, r)
								} else {
									newRarities = append(newRarities, 100) // append a character
								}
							} else if r == 100 { // character
								if charactersCanBeLevelled {
									newRarities = append(newRarities, r)
								} else {
									newRarities = append(newRarities, int64(rand.Intn(3))) // append random rarity Chao
								}
							} else { // should never happen
								panic(fmt.Errorf("invalid rarity '" + strconv.Itoa(int(r)) + "'")) // TODO: use better way to handle
							}
						}
						return newRarities, true
					}
					
					player.ChaoRouletteGroup.ChaoWheelOptions = netobj.DefaultChaoWheelOptions(player.PlayerState) // create a new wheel
					newRarities, ok := fixRarities(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity)
					if !ok { // if player is entirely unable to upgrade anything
						// TODO: this method is super-hacky - ideally we'd want to return an error code for this situation
						player.ChaoRouletteGroup.ChaoWheelOptions.SpinCost = player.PlayerState.NumChaoRouletteTicket + player.PlayerState.NumRedRings // make it impossible for player to use roulette
					} else { // if player can upgrade
						player.ChaoRouletteGroup.ChaoWheelOptions.Rarity = newRarities
					}
					newItems, newRarities, err := roulette.GetRandomChaoRouletteItems(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity, player.GetAllNonMaxedCharacters(), player.GetAllNonMaxedChao())
					if err != nil {
						helper.InternalErr("Error getting new items", err)
						return
					}
					player.ChaoRouletteGroup.WheelChao = newItems
					player.ChaoRouletteGroup.ChaoWheelOptions.Rarity = newRarities
				}
			} else {
				helper.Warn("item '" + wonItem + "' not found")
			}
		}

		helper.DebugOut("Time now: %v", time.Now().UTC().Unix())
		helper.DebugOut("RoulettePeriodEnd: %v", player.RouletteInfo.RoulettePeriodEnd)
		endPeriod := player.RouletteInfo.RoulettePeriodEnd
		helper.DebugOut("Time now (passed): %v", time.Now().UTC().Unix())
		helper.DebugOut("End period (passed): %v", endPeriod)
		if time.Now().UTC().Unix() > endPeriod { // TODO: Do we still need this?
			player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
			helper.DebugOut("New roulette period")
			helper.DebugOut("RouletteCountInPeriod: %v", player.RouletteInfo.RouletteCountInPeriod)
		}

		// generate NEXT! wheel
		if wonItem != strconv.Itoa(enums.IDTypeItemRouletteWin) {
			player.RouletteInfo.RouletteCountInPeriod++ // we've spun an additional time
			// TODO: fix behavior to make it so that a free spin is properly awarded every time when landing on an ItemRouletteWin spot.
			if player.RouletteInfo.RouletteCountInPeriod > freeSpins {
				// we've run out of free spins for the period
				player.PlayerState.NumRouletteTicket--
			}
		}
		numRouletteTicket := player.PlayerState.NumRouletteTicket
		player.OptionUserResult.NumItemRoulette++
		rouletteCount := player.RouletteInfo.RouletteCountInPeriod // get amount of times we've spun the wheel today
		//player.LastWheelOptions = netobj.DefaultWheelOptions(numRouletteTicket, rouletteCount) // create wheel
		//oldRanking := player.LastWheelOptions.RouletteRank
		player.LastWheelOptions = netobj.UpgradeWheelOptions(player.LastWheelOptions, numRouletteTicket, rouletteCount, freeSpins) // create wheel
		if player.RouletteInfo.GotJackpotThisPeriod {
			player.LastWheelOptions.NumJackpotRing = 1
		}
		if wonItem == strconv.Itoa(enums.IDTypeItemRouletteWin) && oldRanking == enums.WheelRankSuper { // won jackpot in super wheel
			helper.DebugOut("Won jackpot in super wheel")
			player.RouletteInfo.GotJackpotThisPeriod = true
		}
	} else {
		// do not modify the wheel, set error status
		responseStatus = status.RouletteUseLimit
	}

	baseInfo := helper.BaseInfo(emess.OK, responseStatus)
	response := responses.WheelSpin(baseInfo, player.PlayerState, player.CharacterState, player.ChaoState, player.LastWheelOptions)
	//response.BaseResponse = responses.NewBaseResponseV(baseInfo, request.Version)
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeSpinItemRoulette)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeSpinItemRoulette)", err)
	}
}
