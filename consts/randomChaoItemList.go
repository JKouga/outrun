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
<<<<<<< HEAD
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
=======
	enums.CTStrSonic:    0.81, // Initial character - all groups
	enums.CTStrTails:    0.81, // Obtained in story mode - all groups
	enums.CTStrKnuckles: 0.81, // Obtained in story mode - all groups
	//enums.CTStrAmy:      0.81, // Group 2
	//enums.CTStrBig:      0.81, // Group 2
	// enums.CTStrBlaze:           0.81, // Group 3
	enums.CTStrCharmy: 0.81, // Group 1
	//enums.CTStrCream:  0.81, // Group 2
	enums.CTStrEspio:  0.81, // Group 1
	//enums.CTStrMephiles:        0.81, // Group 2
	enums.CTStrOmega: 0.81, // Group 1
	//enums.CTStrPSISilver:       0.81, // Group 2
	enums.CTStrRouge:  0.81, // Group 1
	enums.CTStrShadow: 0.81, // Group 1
	// enums.CTStrMarine:          0.81, // Group 3
	// enums.CTStrTangle: 0.81, // Group 3
	// enums.CTStrWhisper: 0.81, // Group 3
	// enums.CTStrSticks:          0.81, // Group 3
	// enums.CTStrTikal:           0.81, // Group 3
	enums.CTStrVector: 0.81, // Group 1
	// enums.CTStrWerehog:         0.81, // Group 3
	//enums.CTStrClassicSonic: 0.81, // Group 2
	enums.CTStrMetalSonic:      0.81, // Group 1
	//enums.CTStrSilver:      0.81, // Group 2

	//enums.CTStrAmitieAmy:       1.81, // Event (Puyo Puyo Quest)
	// enums.CTStrGothicAmy:       3.00, // Revival Event
	// enums.CTStrHalloweenShadow: 1.0, // Event (Halloween)
	// enums.CTStrHalloweenRouge:  1.0, // Event (Halloween)
	//enums.CTStrHalloweenOmega:  1.81, // Event (Halloween)
	// enums.CTStrXMasSonic:       1.0, // Event (Christmas)
	// enums.CTStrXMasTails:       1.81, // Event (Christmas)
	// enums.CTStrXMasKnuckles:    1.0, // Event (Christmas)
	enums.CTStrXT:              1.81, // Revival Event (Christmas)
>>>>>>> ad39a77884de33cfbc226015ecf504a3ae1fdc5a
}

var RandomChaoWheelChaoPrizes = map[string]float64{
	// TODO: Balance these
	//enums.ChaoIDStrHeroChao:             3.0, // Event (Animal Rescue event 1.0) - Group 2
	//enums.ChaoIDStrGoldChao:             3.0, // Event (Animal Rescue event 1.0) - Group 2
	enums.ChaoIDStrDarkChao:             3.02, // Event (Animal Rescue event 1.0) - Group 1
	enums.ChaoIDStrJewelChao:            3.02, // Event (Animal Rescue event 1.0) - Group 1
	// enums.ChaoIDStrNormalChao:           3.7, // Event (Animal Rescue event 1.0) - Group 3
	// enums.ChaoIDStrOmochao:              3.7, // Event (Animal Rescue event 1.0) - Group 3
	//enums.ChaoIDStrRCMonkey:             3.7, // Event (Animal Rescue event 1.0) - Group 2
	enums.ChaoIDStrRCSpring:             3.02, // Permanent Pool
	enums.ChaoIDStrRCElectromagnet:      3.02, // Permanent Pool
	enums.ChaoIDStrBabyCyanWisp:         3.02, // Permanent Pool
	enums.ChaoIDStrBabyIndigoWisp:       3.02, // Permanent Pool
	enums.ChaoIDStrBabyYellowWisp:       3.02, // Permanent Pool
	enums.ChaoIDStrRCPinwheel:           3.02, // Group 1
	enums.ChaoIDStrRCPiggyBank:          3.02, // Permanent Pool
	// enums.ChaoIDStrRCBalloon:            3.7, // Group 3
	enums.ChaoIDStrEasterChao:           3.02, // Event (Easter; Increase Odds During Event) - Group 1
	//enums.ChaoIDStrEasterBunny:          1.98, // Event (Easter; Increase Odds During Event) - Group 2
	//enums.ChaoIDStrMerlina:              1.5, // Event (Easter: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	enums.ChaoIDStrPurplePapurisu:       6.8, // Event (Puyo Puyo Quest)
	//enums.ChaoIDStrSuketoudara:          2.5, // Event (Puyo Puyo Quest)
<<<<<<< HEAD
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
=======
	//enums.ChaoIDStrCarbuncle:            3.67, // Event (Puyo Puyo Quest: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	enums.ChaoIDStrEggChao:           2.26, // Permanent Pool
	enums.ChaoIDStrPumpkinChao:       2.26, // Permanent Pool
	enums.ChaoIDStrSkullChao:         2.26, // Permanent Pool
	enums.ChaoIDStrYacker:            2.26, // Permanent Pool
	enums.ChaoIDStrRCGoldenPiggyBank: 2.26, // Permanent Pool
	enums.ChaoIDStrWizardChao:        2.26, // Permanent Pool
	enums.ChaoIDStrRCTurtle:          2.26, // Group 1
	// enums.ChaoIDStrRCUFO:             1.98, // Group 3
	//enums.ChaoIDStrRCBomber:          1.98, // Group 2
	enums.ChaoIDStrStarShapedMissile:    2.26, // Event (Zazz Raid Boss; Increase Odds During Event) - Group 1
	// enums.ChaoIDStrRCSatellite:          1.98, // Event (Zazz Raid Boss; Increase Odds During Event) - Group 3
	//enums.ChaoIDStrRCMoonMech:           0.0, // Event (Zazz Raid Boss; Only Obtainable through the Raid Boss Roulette, which is currently unavailable)
	// enums.ChaoIDStrRappy:                4.25, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrKuna:                 3.67, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrMagLv1:               6.8, // Event (Phantasy Star Online 2)
	//enums.ChaoIDStrBlowfishTransporter:  1.98, // Event (Tropical Coast; Increase Odds During Event) - Group 2
	//enums.ChaoIDStrMotherWisp:           1.2, // Event (Tropical Coast: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	enums.ChaoIDStrMarineChao:           2.26, // Event (Tropical Coast; Increase Odds During Event) - Group 1
	//enums.ChaoIDStrGenesis:              1.5, // Event (Birthday: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	// enums.ChaoIDStrCartridge:            1.98, // Event (Birthday; Increase Odds During Event) - Group 3
 	enums.ChaoIDStrDeathEgg:             1.47, // Event (Birthday; Increase Odds During Event) - Group 1
	//enums.ChaoIDStrRCFighter:            1.98, // Group 2
	enums.ChaoIDStrRCHovercraft:         2.26, // Group 1
	// enums.ChaoIDStrRCHelicopter:         1.98, // Group 3
	//enums.ChaoIDStrGreenCrystalMonsterS: 1.98, // Group 2
	enums.ChaoIDStrGreenCrystalMonsterL: 2.26, // Group 1
	// enums.ChaoIDStrRCAirship:            1.98, // Group 3
	//enums.ChaoIDStrMagicLamp:            1.98, // Event (Desert Ruins and Animal Rescue 2.0; Increase Odds During Event) - Group 2
	enums.ChaoIDStrDesertChao:           2.26, // Event (Desert Ruins; Increase Odds During Event) - Group 1
	//enums.ChaoIDStrErazorDjinn:          2.5, // Event (Desert Ruins: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	//enums.ChaoIDStrNightopian:           4.25, // Event (NiGHTS)
>>>>>>> ad39a77884de33cfbc226015ecf504a3ae1fdc5a
	//enums.ChaoIDStrNiGHTS:               2.0, // Event (NiGHTS)
	// enums.ChaoIDStrReala:                3.4, // Event (NiGHTS)
	// enums.ChaoIDStrSonicOmochao:         1.98, // Event (Team Sonic Omochao) - Group 3
	//enums.ChaoIDStrTailsOmochao:         1.98, // Event (Team Sonic Omochao) - Group 2
	enums.ChaoIDStrKnucklesOmochao:      2.26, // Event (Team Sonic Omochao) - Group 1
	//enums.ChaoIDStrKingBoomBoo:          1.5, // Event (Halloween: Premium Roulette for Timed Mode Event Only; Obtainable through Rewards List Only for Story Event)
	// enums.ChaoIDStrBoo:                  1.98, // Event (Halloween; Increase Odds During Event) - Group 3
	//enums.ChaoIDStrHalloweenChao:        1.98, // Event (Halloween; Increase Odds During Event) - Group 2
	//enums.ChaoIDStrHeavyBomb:            1.2, // Event (Fantasy Zone)
	// enums.ChaoIDStrOPapa:                3.4, // Event (Fantasy Zone)
	enums.ChaoIDStrOpaOpa:               3.67, // Event (Fantasy Zone)
	enums.ChaoIDStrBlockBomb:  2.26, // Group 1
	// enums.ChaoIDStrHunkofMeat: 1.98, // Group 3
	// enums.ChaoIDStrYeti:                 1.5, // Event (Christmas)
	// enums.ChaoIDStrSnowChao:             2.0, // Event (Christmas)
	// enums.ChaoIDStrChristmasYeti:        1.5, // Event (Christmas)
	// enums.ChaoIDStrChristmasNiGHTS:      1.5, // Event (Christmas NiGHTS)
	// enums.ChaoIDStrIdeya:                1.5, // Event (Christmas NiGHTS)
	// enums.ChaoIDStrChristmasNightopian:  1.5, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrOrbot:      1.98, // Group 2
	enums.ChaoIDStrCubot:      2.26, // Group 1
	enums.ChaoIDStrLightChaos: 1.47, // Permanent Pool
	enums.ChaoIDStrHeroChaos:  1.47, // Permanent Pool
	enums.ChaoIDStrDarkChaos:  1.47, // Permanent Pool
	enums.ChaoIDStrChip:       1.47, // Permanent Pool
	//enums.ChaoIDStrShahra:               4.0, // Runners' League Story Mode
	enums.ChaoIDStrCaliburn:         1.47, // Permanent Pool
	enums.ChaoIDStrKingArthursGhost: 1.47, // Permanent Pool
	//enums.ChaoIDStrRCTornado:        1.47, // Group 2
	enums.ChaoIDStrRCBattleCruiser:  1.47, // Group 1
	// enums.ChaoIDStrRedCrystalMonsterS: 1.38, // Group 3
	//enums.ChaoIDStrRedCrystalMonsterL: 1.47, // Group 2
	enums.ChaoIDStrGoldenGoose:        1.47, // Group 1
	//enums.ChaoIDStrRCPirateSpaceship: 1.47, // Group 2
	// enums.ChaoIDStrGoldenAngel:       1.38, // Group 3
	enums.ChaoIDStrRCTornado2:           1.47, // Event (Sonic Adventure; Increase Odds During Event) - Group 1
	// enums.ChaoIDStrChaos:                1.38, // Event (Sonic Adventure; Increase Odds During Event) - Group 3
	// enums.ChaoIDStrOrca:                 1.98, // Event (Sonic Adventure; Increase Odds During Event) - Group 3
	//enums.ChaoIDStrChaoWalker:           0.0, // Daily Battle
	enums.ChaoIDStrDarkQueen:            3.67, // Runners' League Timed Mode
	//enums.ChaoIDStrRCBlockFace: 1.47, // Group 2
	//enums.ChaoIDStrDFekt:                0.0, // Revival Event (assets TBD)
<<<<<<< HEAD
	//enums.ChaoIDStrDarkChaoWalker:       0.0, // Daily Battle?
	//enums.ChaoIDStrPrideChaoL:           1.5,
	//enums.ChaoIDStrPrideChaoG:           1.5,
	//enums.ChaoIDStrPrideChaoB:           1.5,
	//enums.ChaoIDStrPrideChaoT:           1.5,
	//enums.ChaoIDStrPrideChaoP:           1.5,
	//enums.ChaoIDStrPrideChaoA:           1.5,
	//enums.ChaoIDStrPrideChaoNB:          1.5,
=======
	// enums.ChaoIDStrDarkChaoWalker:       1.38, // Daily Battle? - Group 3
	//enums.ChaoIDStrPrideChaoL:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoG:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoB:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoT:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoP:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoA:           3.09, // Revival Event (Pride Month Celebration)
	//enums.ChaoIDStrPrideChaoNB:          3.09, // Revival Event (Pride Month Celebration)
>>>>>>> ad39a77884de33cfbc226015ecf504a3ae1fdc5a
}
