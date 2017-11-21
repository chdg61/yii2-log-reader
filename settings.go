package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Settings struct {
	RegexpCheck string `json:"regexp_check"`
	RegexpHeader string `json:"regexp_check"`
}

var settings = Settings{}

const fileSettings = "settings.json";

func init() {
	if _, err := os.Stat(fileSettings); !os.IsNotExist(err) {
		file, _ := ioutil.ReadFile(fileSettings)
		fmt.Println(file)
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
