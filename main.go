//go:generate fileb0x b0x.yml
package main

import (
	"github.com/ImVexed/muon"
	"net/http"
)

func main() {
	fileHandler := http.FileServer(http.Dir("assets"))

	cfg := &muon.Config{
		Title:      "Hello, World!",
		Height:     540,
		Width:      1024,
		Titled:     true,
		Resizeable: true,
	}

	m := muon.New(cfg, fileHandler)

	m.Bind("add", add)

	if err := m.Start(); err != nil {
		panic(err)
	}
}

func add(a float64, b float64) float64 {
	return a + b
}
