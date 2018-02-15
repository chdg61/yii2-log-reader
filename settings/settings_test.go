package settings

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestInitSetting(t *testing.T) {

	t.Run("TestDefaultFileName", func(t *testing.T) {
		if(fileSettings != "settings.json"){
			t.Error("Not validate default file name setting")
		}
	})

	t.Run("TestDefaultSettings", func(t *testing.T) {
		if(settings.RegexpCheck != "\\d{4}-\\d{1,2}-\\d{1,2}\\s[\\d:]*\\s\\[.*?\\]"){
			t.Error("Not valid default Settings")
		}
	})

	file, _ := ioutil.TempFile(os.TempDir(), "test__")
	defer os.Remove(file.Name())

	file.WriteString(`{"regexp_check": "test"}`)
	fileSettings = file.Name()
	initSetting();
	if(settings.RegexpCheck != "test"){
		t.Error("Not rechange setting")
	}

}

func TestGetInstant(t *testing.T) {
	settings = Settings{}

	if(settings != GetInstant()){
		t.Error("Not valid singlton")
	}
}
