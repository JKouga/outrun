package netobj

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/Mtbcooler/outrun/config"
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/enums"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/jinzhu/now"
)

type WheelOptions struct {
	Items                []string   `json:"items"`
	Item                 []int64    `json:"item"`
	ItemWeight           []int64    `json:"itemWeight"`
	ItemWon              int64      `json:"itemWon"`
	NextFreeSpin         int64      `json:"nextFreeSpin"` // midnight (start of next day)
	SpinCost             int64      `json:"spinCost"`
	RouletteRank         int64      `json:"rouletteRank"`
	NumRouletteToken     int64      `json:"numRouletteToken"`
	NumJackpotRing       int64      `json:"numJackpotRing"`
	NumRemainingRoulette int64      `json:"numRemainingRoulette"`
	ItemList             []obj.Item `json:"itemList"`
}

func DefaultWheelOptions(numRouletteTicket, rouletteCountInPeriod, rouletteRank int64) WheelOptions {
	// TODO: Modifying this seems like a good way of figuring out what the game thinks each ID means in terms of items.
	// const the below
	// NOTE: Free spins occur when numRemainingRoulette > numRouletteToken
	//items := []string{"200000", "120000", "120001", "120002", "200000", "900000", "120003", "120004"}
	rouletteGenMode := rand.Intn(3)
	items := []string{strconv.Itoa(enums.IDTypeItemRouletteWin)} // first item is always jackpot/big/super
	item := []int64{1}
	// There are currently three roulette generation modes:
	// Mode 0: Classic mode
	// Mode 1: Vertical dual win (based off a pattern in the OG server)
	// Mode 2: Classic mode but with two win spots placed horizontally instead of one win spot on the top
	itemWeight := []int64{1250, 1250, 1250, 1250, 1250, 1250, 1250, 1250}
	switch rouletteGenMode {
	case 1:
		randomItem1 := consts.RandomItemListNormalWheel[rand.Intn(len(consts.RandomItemListNormalWheel))]
		randomItemAmount1 := consts.NormalWheelItemAmountRange[randomItem1].GetRandom()
		randomItem2 := consts.RandomItemListNormalWheel[rand.Intn(len(consts.RandomItemListNormalWheel))]
		randomItemAmount2 := consts.NormalWheelItemAmountRange[randomItem2].GetRandom()
		items = append(items, randomItem1)
		item = append(item, randomItemAmount1)
		items = append(items, randomItem2)
		item = append(item, randomItemAmount2)
		items = append(items, randomItem1)
		item = append(item, randomItemAmount1)
		if rouletteRank != enums.WheelRankSuper {
			items = append(items, strconv.Itoa(enums.IDTypeItemRouletteWin))
			item = append(item, 1)
		} else {
			items = append(items, randomItem2)
			item = append(item, randomItemAmount2)
		}
		items = append(items, randomItem1)
		item = append(item, randomItemAmount1)
		items = append(items, randomItem2)
		item = append(item, randomItemAmount2)
		items = append(items, randomItem1)
		item = append(item, randomItemAmount1)
	default:
		for _ = range make([]byte, 7) { // loop 7 times
			randomItem := consts.RandomItemListNormalWheel[rand.Intn(len(consts.RandomItemListNormalWheel))]
			randomItemAmount := consts.NormalWheelItemAmountRange[randomItem].GetRandom()
			items = append(items, randomItem)
			item = append(item, randomItemAmount)
		}
		if rouletteGenMode == 2 && rouletteRank != enums.WheelRankSuper {
			randomItem := consts.RandomItemListNormalWheel[rand.Intn(len(consts.RandomItemListNormalWheel))]
			randomItemAmount := consts.NormalWheelItemAmountRange[randomItem].GetRandom()
			items[0] = randomItem
			item[0] = randomItemAmount
			items[2] = strconv.Itoa(enums.IDTypeItemRouletteWin)
			item[2] = 1
			items[6] = strconv.Itoa(enums.IDTypeItemRouletteWin)
			item[6] = 1
		}
	}
	//itemWon := int64(0)
	itemWon := int64(rand.Intn(len(items)))   //TODO: adjust this to accurately represent item weights
	nextFreeSpin := now.EndOfDay().Unix() + 1 // midnight
	spinCost := int64(15)
	//rouletteRank := int64(enums.WheelRankNormal)
	//numRouletteToken := playerState.NumRouletteTicket
	numRouletteToken := numRouletteTicket // The game uses the _current_ value, not as if it was in the past (This is hard to explain, maybe TODO: explain this better?)
	numJackpotRing := int64(consts.RouletteJackpotRings)
	// TODO: get rid of logic here!
	numRemainingRoulette := numRouletteToken + consts.RouletteFreeSpins - rouletteCountInPeriod // TODO: is this proper?
	if numRemainingRoulette < numRouletteToken {
		numRemainingRoulette = numRouletteToken
	}
	itemList := []obj.Item{}
	out := WheelOptions{
		items,
		item,
		itemWeight,
		itemWon,
		nextFreeSpin,
		spinCost,
		rouletteRank,
		numRouletteToken,
		numJackpotRing,
		numRemainingRoulette,
		itemList,
	}
	return out
}

func UpgradeWheelOptions(origWheel WheelOptions, numRouletteTicket, rouletteCountInPeriod int64) WheelOptions {
	rouletteRank := origWheel.RouletteRank
	if origWheel.Items[origWheel.ItemWon] == strconv.Itoa(enums.IDTypeItemRouletteWin) { // if landed on big/super or jackpot
		landedOnUpgrade := origWheel.RouletteRank == enums.WheelRankNormal || origWheel.RouletteRank == enums.WheelRankBig
		if config.CFile.DebugPrints {
			log.Printf("%v\n", origWheel.RouletteRank)
			log.Printf("%v\n", landedOnUpgrade)
		}
		if landedOnUpgrade {
			if config.CFile.DebugPrints {
				log.Println("landedOnUpgrade")
			}
			rouletteRank++ // increase the rank
		} else {
			if config.CFile.DebugPrints {
				log.Println("NOT landedOnUpgrade")
			}
			rouletteRank = enums.WheelRankNormal
		}
	} else {
		rouletteRank = enums.WheelRankNormal
	}
	newWheel := DefaultWheelOptions(numRouletteTicket, rouletteCountInPeriod, rouletteRank)
	return newWheel
}
