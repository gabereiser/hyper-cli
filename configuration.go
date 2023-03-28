/*
 * Copyright (c) 2023 Red Engine Games, LLC.
 * All Rights Reserved.
 */

package hyper_cli

import (
	"encoding/json"
	"os"
)

type HyperConfiguration struct {
	Prefix    string
	InputDir  string
	OutputDir string
	Theme     string
	DevPort   int
}

func LoadConfiguration() HyperConfiguration {
	file, _ := os.ReadFile("hyper.config.json")
	config := DefaultConfiguration()
	_ = json.Unmarshal(file, &config)
	return config
}

func DefaultConfiguration() HyperConfiguration {
	return HyperConfiguration{
		Prefix:    "",
		InputDir:  ".",
		OutputDir: "dist",
		Theme:     "default",
		DevPort:   8080,
	}
}
