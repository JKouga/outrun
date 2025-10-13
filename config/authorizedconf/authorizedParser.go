package authorizedconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Defaults = map[string]interface{}{
	"DAuthorizedIDs":           []string{},
	"DAllowCustomBlockMessage": false,
	"DCustomBlockMessage":      "",
}

type ConfigFile struct {
	AuthorizedIDs           []string `json:"ids,omitempty"`
	AllowCustomBlockMessage bool     `json:"useCustomBlockMessage,omitempty"`
	CustomBlockMessage      string   `json:"customBlockMessage,omitempty"`
}

var CFile ConfigFile

func Parse(filename string) error {
	CFile = ConfigFile{
		Defaults["DAuthorizedIDs"].([]string),
		Defaults["DAllowCustomBlockMessage"].(bool),
		Defaults["DCustomBlockMessage"].(string),
	}
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &CFile)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}