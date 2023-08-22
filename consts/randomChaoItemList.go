package consts

import "github.com/RunnersRevival/outrun/enums"

type PrizeInfo struct {
	AppearanceChance float64 // % chance for it to be chosen to be in wheel by the server
	Type             int64   // 0 for Chao, 1 for Character
}

// A 'load' as depicted below is the chance for the server to pick
// the associated item, where chosen is if randFloat(0, 100) < load.
// IMPORTANT: This load is exclusive to the rarity of the Chao that
// is being chosen by the server.

var RandomChaoWheelCharacterPrizes = map[string]float64{
	// characterID: load
	// Hopefully this should sum up to 100 just for
	// simplicity, but it shouldn't be a requirement.
	enums.CTStrSonic:    0.3, // Initial character
	enums.CTStrTails:    0.3, // Obtained in story mode
	enums.CTStrKnuckles: 0.3, // Obtained in story mode
	enums.CTStrAmy:      0.3,
	enums.CTStrBig:      0.3,
	//enums.CTStrBlaze:           1.0, // Revival Event (Sonic Rush)
	enums.CTStrCharmy: 0.3,
	enums.CTStrCream:  0.3,
	enums.CTStrEspio:  0.3,
	//enums.CTStrMephiles:        0.0, // Revival Event
	enums.CTStrOmega: 0.3,
	//enums.CTStrPSISilver:       0.0, // Revival Event
	enums.CTStrRouge:  0.3,
	enums.CTStrShadow: 0.3,
	enums.CTStrMarine:          1.0, // Revival Event
	enums.CTStrTangle: 1.0, // Revival Event
	enums.CTStrWhisper: 1.0, // Revival Event
	//enums.CTStrSticks:          0.3, // Revival Event
	//enums.CTStrTikal:           1.5, // Event (Sonic Adventure)
	enums.CTStrVector: 0.3,
	//enums.CTStrWerehog:         1.5, // Revival Event
	//enums.CTStrClassicSonic: 1.0, // Event (Birthday)
	//enums.CTStrMetalSonic:      0.0, // Revival Event

	// The below characters shouldn't be activated until event characters are fixed!
	enums.CTStrAmitieAmy:       1.0, // Event (Puyo Puyo Quest)
	//enums.CTStrGothicAmy:       0.0, // Not given out on Revival
	//enums.CTStrHalloweenShadow: 1.0,   // Event (Halloween)
	//enums.CTStrHalloweenRouge:  1.0,   // Event (Halloween)
	//enums.CTStrHalloweenOmega:  1.0,   // Event (Halloween)
	//enums.CTStrXMasSonic:       1.0, // Event (Christmas)
	//enums.CTStrXMasTails:       1.0, // Event (Christmas)
	//enums.CTStrXMasKnuckles:    1.0, // Event (Christmas)
}

var RandomChaoWheelChaoPrizes = map[string]float64{
	// TODO: Balance these
	enums.ChaoIDStrHeroChao:             2.2, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrGoldChao:             2.2, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrDarkChao:             2.2, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrJewelChao:            2.2, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrNormalChao:           3.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrOmochao:              3.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrRCMonkey:             1.0, // Event (Animal Rescue event 1.0)
	enums.ChaoIDStrRCSpring:             2.0,
	enums.ChaoIDStrRCElectromagnet:      2.0,
	enums.ChaoIDStrBabyCyanWisp:         3.0,
	enums.ChaoIDStrBabyIndigoWisp:       3.0,
	enums.ChaoIDStrBabyYellowWisp:       3.0,
	enums.ChaoIDStrRCPinwheel:           2.0,
	enums.ChaoIDStrRCPiggyBank:          1.2,
	enums.ChaoIDStrRCBalloon:            2.0,
	//enums.ChaoIDStrEasterChao:           1.0, // Event (Easter)
	//enums.ChaoIDStrEasterBunny:          1.0, // Event (Easter)
	//enums.ChaoIDStrMerlina:              1.5, // Event (Easter)
	//enums.ChaoIDStrPurplePapurisu:       2.5, // Event (Puyo Puyo Quest)
	//enums.ChaoIDStrSuketoudara:          2.5, // Event (Puyo Puyo Quest)
	//enums.ChaoIDStrCarbuncle:            2.5, // Event (Puyo Puyo Quest)
	enums.ChaoIDStrEggChao:           1.0,
	enums.ChaoIDStrPumpkinChao:       1.0,
	enums.ChaoIDStrSkullChao:         1.1,
	enums.ChaoIDStrYacker:            1.0,
	enums.ChaoIDStrRCGoldenPiggyBank: 1.0,
	enums.ChaoIDStrWizardChao:        1.0,
	enums.ChaoIDStrRCTurtle:          1.1,
	enums.ChaoIDStrRCUFO:             1.1,
	enums.ChaoIDStrRCBomber:          1.0,
	//enums.ChaoIDStrStarShapedMissile:    0.0, // Event (Zazz Raid Boss)
	//enums.ChaoIDStrRCSatellite:          0.0, // Event (Zazz Raid Boss)
	//enums.ChaoIDStrRCMoonMech:           0.0, // Event (Zazz Raid Boss?)
	//enums.ChaoIDStrRappy:                1.0, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrKuna:                 1.0, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrMagLv1:               1.0, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrBlowfishTransporter:  1.5, // Event (Tropical Coast)
	//enums.ChaoIDStrMotherWisp:           1.2, // Event (Tropical Coast)
	//enums.ChaoIDStrMarineChao:           1.5, // Event (Tropical Coast)
	//enums.ChaoIDStrGenesis:              1.5, // Event (Birthday)
	//enums.ChaoIDStrCartridge:            1.5, // Event (Birthday)
	//enums.ChaoIDStrDeathEgg:             1.0, // Event (Birthday)
	enums.ChaoIDStrRCFighter:            1.0,
	enums.ChaoIDStrRCHovercraft:         1.0,
	enums.ChaoIDStrRCHelicopter:         1.0,
	enums.ChaoIDStrGreenCrystalMonsterS: 1.0,
	enums.ChaoIDStrGreenCrystalMonsterL: 1.0,
	enums.ChaoIDStrRCAirship:            1.0,
	enums.ChaoIDStrMagicLamp:            2.0, // Event (Desert Ruins and Animal Rescue 2.0)
	//enums.ChaoIDStrDesertChao:           3.0, // Event (Desert Ruins)
	//enums.ChaoIDStrErazorDjinn:          2.5, // Event (Desert Ruins)
	//enums.ChaoIDStrNightopian:           2.0, // Event (NiGHTS)
	//enums.ChaoIDStrNiGHTS:               2.0, // Event (NiGHTS)
	//enums.ChaoIDStrReala:                2.0, // Event (NiGHTS)
	//enums.ChaoIDStrSonicOmochao:         0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrTailsOmochao:         0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrKnucklesOmochao:      0.0, // Event (Team Sonic Omochao)
	//enums.ChaoIDStrKingBoomBoo:          1.5, // Event (Halloween)
	//enums.ChaoIDStrBoo:                  1.5, // Event (Halloween)
	//enums.ChaoIDStrHalloweenChao:        1.5, // Event (Halloween)
	//enums.ChaoIDStrHeavyBomb:            1.2, // Event (Fantasy Zone)
	//enums.ChaoIDStrOPapa:                1.5, // Event (Fantasy Zone)
	//enums.ChaoIDStrOpaOpa:               1.5, // Event (Fantasy Zone)
	enums.ChaoIDStrBlockBomb:  1.0,
	enums.ChaoIDStrHunkofMeat: 1.3,
	//enums.ChaoIDStrYeti:                 1.5, // Event (Christmas)
	//enums.ChaoIDStrSnowChao:             2.0, // Event (Christmas)
	//enums.ChaoIDStrChristmasYeti:        1.5, // Event (Christmas)
	//enums.ChaoIDStrChristmasNiGHTS:      1.5, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrIdeya:                1.5, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrChristmasNightopian:  1.5, // Event (Christmas NiGHTS)
	enums.ChaoIDStrOrbot:      1.0,
	enums.ChaoIDStrCubot:      1.0,
	enums.ChaoIDStrLightChaos: 1.5,
	enums.ChaoIDStrHeroChaos:  1.5,
	enums.ChaoIDStrDarkChaos:  1.5,
	enums.ChaoIDStrChip:       1.5,
	//enums.ChaoIDStrShahra:               0.0, // Runners' League
	enums.ChaoIDStrCaliburn:         1.2,
	enums.ChaoIDStrKingArthursGhost: 1.0,
	enums.ChaoIDStrRCTornado:        1.0,
	enums.ChaoIDStrRCBattleCruiser:  0.7,
	enums.ChaoIDStrRedCrystalMonsterS: 1.0,
	enums.ChaoIDStrRedCrystalMonsterL: 1.0,
	enums.ChaoIDStrGoldenGoose:        1.0,
	enums.ChaoIDStrRCPirateSpaceship: 0.7,
	enums.ChaoIDStrGoldenAngel:       1.0,
	//enums.ChaoIDStrRCTornado2:           1.0, // Event (Sonic Adventure)
	//enums.ChaoIDStrChaos:                1.0, // Event (Sonic Adventure)
	//enums.ChaoIDStrOrca:                 1.5, // Event (Sonic Adventure)
	//enums.ChaoIDStrChaoWalker:           0.0, // Daily Battle
	//enums.ChaoIDStrDarkQueen:            0.0, // Runners' League
	enums.ChaoIDStrRCBlockFace: 1.0,
	//enums.ChaoIDStrDFekt:                0.0, // Revival Event (assets TBD)
	//enums.ChaoIDStrDarkChaoWalker:       0.0, // Daily Battle?
	//enums.ChaoIDStrPrideChaoL:           1.5,
	//enums.ChaoIDStrPrideChaoG:           1.5,
	//enums.ChaoIDStrPrideChaoB:           1.5,
	//enums.ChaoIDStrPrideChaoT:           1.5,
	//enums.ChaoIDStrPrideChaoP:           1.5,
	//enums.ChaoIDStrPrideChaoA:           1.5,
	//enums.ChaoIDStrPrideChaoNB:          1.5,
}
