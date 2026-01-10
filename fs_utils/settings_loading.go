package fs_utils

import (
	"encoding/json"
	"os"
)

func WriteSettings[T any](path string, settings T) error {
	content, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	return os.WriteFile(path, content, 0644)
}

func ReadSettings[T any](path string) (*T, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var settings T
	err = json.Unmarshal(content, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func ReadSettingsOrWriteAndReturnDefault[T any](path string, default_settings T) (*T, error) {
	settings, err := ReadSettings[T](path)
	if err != nil {
		err = WriteSettings(path, default_settings)
		if err != nil {
			return nil, err
		}

		return &default_settings, nil
	}

	return settings, nil
}
