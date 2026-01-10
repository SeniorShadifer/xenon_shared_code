package fs_utils_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/SeniorShadifer/xenon_shared_code/fs_utils"
)

type TestSettingsStruct struct{ Message string }

func TestReadSettings(t *testing.T) {
	settings := TestSettingsStruct{Message: "Hello, world!"}
	path := "test_settings_file.json"

	serialized_settings, err := json.Marshal(settings)
	if err != nil {
		t.Fatal("Failed to serialize settings.")
	}

	err = os.WriteFile(path, serialized_settings, 0644)
	if err != nil {
		t.Fatal("Failed to write settings to file.")
	}

	readed_settings, err := fs_utils.ReadSettings[TestSettingsStruct](path)
	if err != nil {
		t.Fatal("Cannot read settings:", err)
	}

	if !reflect.DeepEqual(&settings, readed_settings) {
		t.Fatal("Settings", readed_settings, "not equals to expected", settings)
	}
}
