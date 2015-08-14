package main

// import "github.com/giantswarm/generic-types-go"

type App struct {
	Name   string            `json:"name"`
	Image  string            `json:"image"`
	Port   string            `json:"port"`
	Domain string            `json:"domain"`
	Env    map[string]string `json:"env"`
	Args   []string          `json:"args"`
}
