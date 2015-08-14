package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type App struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Port  int      `json:"port"`
	Env   []string `json:"env"`
	Args  []string `json:"args"`
}

func main() {
	appJSON, err := ioutil.ReadFile("./app.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	var appGo App
	appGo, err = validateAppDefinitionFile(appJSON, appGo)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%+v", appGo)
	// parseAppDefinitionToUnitFile()
	// writeUnitFile()
}

func validateAppDefinitionFile(appJSON []byte, appGo App) (App, error) {
	err := json.Unmarshal(appJSON, &appGo)
	if err != nil {
		return nil, err
	}
	return appGo, nil
}

// func parseAppDefinitionToUnitFile() {

// }

// func writeUnitFile() {

// }
