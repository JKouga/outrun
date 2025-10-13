# Outrun for Revival

**THIS BRANCH IS NOT INTENDED TO BE SELF-HOSTED. This version of Outrun for Revival is only designed to work with 2.2.0. For a version of Outrun for Revival which is self-hostable, go [here](https://github.com/RunnersRevival/outrun/tree/self-hostable) instead.**

### Summary

Outrun for Revival is a fork of Outrun, a custom server for Sonic Runners reverse engineered from the [Sonic Runners Revival](https://sonicrunners.com) project back during the Open Beta. It is intended for use on the Sonic Runners Revival server, but can be used for your own private servers as well.

### Current functionality

Notable:
  - Timed Mode
  - Story Mode
  - Ring/Red Star Ring keeping
  - Functional shop
  - Character/Chao equipping
  - Character leveling and progression
  - Item/Chao roulette functionality
  - Events
  - Campaigns
  - Basic ranking
  - Login Bonuses
  - Daily Challenge
  - Daily Battles

Functional:
  - Android and iOS support
  - High score keeping
  - In game notices
  - Deep configuration options
  - Powerful RPC control functions
  - Ticker notices
  - Low CPU usage
  - Analytics support
  - Revive Token keeping

### Building

1. [Download and install Go 1.18](https://golang.org/dl/)
2. [Download and install Git](https://git-scm.com/downloads) (for `go get`)
3. Set your [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) environment variable (on Windows, the default is `C:\Users\UsernameHere\go`, while on Linux the default is `/home/usernamehere/go` - only change this if you want to put your work in another directory)
4. Open a terminal/command prompt
5. Use `cd` ([Windows,](https://www.digitalcitizen.life/command-prompt-how-use-basic-commands) [Linux/macOS](https://www.macworld.com/article/2042378/master-the-command-line-navigating-files-and-folders.html)) to navigate to a directory of choice
6. Run `go get github.com/RunnersRevival/outrun` and wait until the command line returns
7. Run `go build github.com/RunnersRevival/outrun` and wait until the build is complete
8. Run the produced executable (`outrun.exe` on Windows, `outrun` on Linux/macOS)

The latest binaries for **Outrun for Revival** can be found [in the releases tab.](https://github.com/RunnersRevival/outrun/releases)

#### Notice regarding self-hosting

This branch of the Outrun for Revival code is not designed to be self-hosted. You will not be able to use 2.0.3 with this codebase, and as such we do not provide patching instructions for iOS and Android in this branch.

### Misc.

Any pull requests deemed code improvements are strongly encouraged. Refactors may be merged into a different branch.

### Credits

Much thanks to:
  - **YPwn**, whose closest point of online social contact I do not know, for creating and running the Sonic Runners Revival server upon which this project based much of its code upon.
  - **[@Sazpaimon](https://github.com/Sazpaimon)** for finding the encryption key I so desparately looked for but could not on my own.
  - **nacabaro** (nacabaro#2138 on Discord) for traffic logging and the discovery of **[DaGuAr](https://www.youtube.com/user/Gorila5)**'s asset archive.

#### Additional assistance
  - Story Mode items
    - lukaafx (Discord @Kalu04#3243)
    - [TemmieFlakes](https://twitter.com/pictochat3)
    - SuperSonic893YT
