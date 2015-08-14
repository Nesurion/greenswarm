package main

import (
	"encoding/json"
	"fmt"
	// "github.com/giantswarm/generic-types-go"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func main() {

	var cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "create new unitfiles for an app.json",
		Long:  `todo: add description`,
		Run: func(cmd *cobra.Command, args []string) {
			generateUnitFiles(parseJSON())
		},
	}

	var cmdStatus = &cobra.Command{
		Use:   "status [app-name]",
		Short: "show status of an app",
		Long:  `todo: add description`,
		Run: func(cmd *cobra.Command, args []string) {
			listUnits(args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: "eugene"}
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdStatus)
	rootCmd.Execute()
}

func parseJSON() App {
	var app App
	appJSON, err := ioutil.ReadFile("./app.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	app, err = validateAppDefinitionFile(appJSON, app)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	return app
}

func validateAppDefinitionFile(data []byte, app App) (App, error) {
	err := json.Unmarshal(data, &app)
	if err != nil {
		return app, err
	}
	return app, nil
}

func generateUnitFiles(app App) {
	tmpl_app := template.Must(template.ParseFiles("./tmpl/app.tmpl"))
	tmpl_sk := template.Must(template.ParseFiles("./tmpl/sk.tmpl"))
	err := tmpl_app.Execute(os.Stdout, app)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = tmpl_sk.Execute(os.Stdout, app)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func listUnits(appname string) {
	// fmt.Sprintf("%s", appname)
	out, err := exec.Command("/usr/local/bin/fleetctl", "list-units").Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%s", out)
}
