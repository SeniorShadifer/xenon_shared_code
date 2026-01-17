package fs_utils_test

import (
	"reflect"
	"testing"

	"github.com/SeniorShadifer/xenon_shared_code/fs_utils"
)

type TestSettingsStruct struct{ Message string }

const path = "test_settings_file.json"

func TestReadAndWriteSettingsFunctions(t *testing.T) {
	settings := TestSettingsStruct{Message: "Hello, world!"}

	err := fs_utils.WriteSettings(path, settings)
	if err != nil {
		t.Fatal("Cannot write settings to file:", err)
	}

	readed_settings, err := fs_utils.ReadSettings[TestSettingsStruct](path)
	if err != nil {
		t.Fatal("Cannot read settings:", err)
	}

	if !reflect.DeepEqual(&settings, readed_settings) {
		t.Fatal("Settings", readed_settings, "not equals to expected", settings)
	}
}

// TODO: Upgrade test
func TestReadSettingsOrWriteAndReturnDefaultFunction(t *testing.T) {
	settings := TestSettingsStruct{Message: "Hello, world!"}

	readed_settings, err := fs_utils.ReadSettingsOrWriteAndReturnDefault(path, settings)
	if err != nil {
		t.Fatal("Failed:", err)
	}

	if !reflect.DeepEqual(&settings, readed_settings) {
		t.Fatal("Settings", readed_settings, "not equals to expected", settings)
	}
}
