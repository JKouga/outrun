package constobjs

import (
    "os/exec"

    "github.com/Mtbcooler/outrun/obj"
)

var DefaultInformations = []obj.Information{
    obj.NewInformation(531, 1, 1465808400, 1466413199, "3__90000001_14"),
    //obj.NewInformation(6001070, 1, 1464336180, 1580608922, "2_[ffff00]'You feel it too, don't you?'[ffffff]\nWelcome to [ff0000]OUTRUN[ffffff], a custom Sonic Runners server!\n\n(Last local Git update: "+getLastGitCommitDate()+")\r\n_10600001_0"),
    obj.NewInformation(6001070, 1, 1464336180, 1580608922, "2_[ffff00]'You feel it too, don't you?'[ffffff]\nWelcome to [ff0000]OUTRUN[ffffff], a custom Sonic Runners server!\r\n_10600001_0"),
    //obj.NewInformation(1000230, 3, 1465981200, 1466413199, "1__90000001_1"),
    //obj.NewInformation(1000157, 600, 1448614800, 1609459199, "1__90000002_1"),
}

func getLastGitCommitDate() string {
    // TODO: remove, since most people likely will not be running Outrun inside of a Git directory
    output, err := exec.Command("git", "log", "-1", "--format=%cd").Output()
    if err != nil {
        return "Unknown"
    }
    return string(output)
}
