package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type Settings struct {
	RegexpCheck string `json:"regexp_check"`
	RegexpHeader string `json:"regexp_header"`
}

var settings = Settings{}

var fileSettings = "settings.json";

func init() {
	initSetting()
}

func initSetting(){
	if _, err := os.Stat(fileSettings); !os.IsNotExist(err) {
		file, _ := ioutil.ReadFile(fileSettings)
		json.Unmarshal(file, &settings)
	}

	if settings.RegexpCheck == "" {
		settings.RegexpCheck = "\\d{4}-\\d{1,2}-\\d{1,2}\\s[\\d:]*\\s\\[.*?\\]";
		settings.RegexpHeader = "(\\d{4}-\\d{1,2}-\\d{1,2}\\s[\\d:]*)\\s\\[(.*?)\\]\\[(.*?)\\]\\[(.*?)\\]\\[(.*?)\\]\\[(.*?)\\]\\s(.*)";
	}
}

func GetInstant() Settings {
	return settings
}
