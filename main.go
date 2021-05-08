//go:generate fileb0x b0x.yml
package main

import (
	"fmt"
	"github.com/ImVexed/muon"
	"github.com/sqweek/dialog"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	dialog.Message("%s", "It developed by akakou.").Title("UREE").Info()
	fileHandler := http.FileServer(http.Dir("assets"))

	cfg := &muon.Config{
		Title:      "Hello, World!",
		Height:     540,
		Width:      1024,
		Titled:     true,
		Resizeable: true,
	}

	m := muon.New(cfg, fileHandler)

	m.Bind("openFile", openFile)
	m.Bind("saveFile", saveFile)
	m.Bind("saveFileAsNew", saveFileAsNew)

	if err := m.Start(); err != nil {
		panic(err)
	}
}

func openFile() string {
	fileName, err := dialog.File().Title("Open").Filter("All Files", "*").Load()
	if err != nil {
		dialog.Message("%s", err).Title("Error").Info()
		return ""
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func saveFile() {
}

func saveFileAsNew(stringBody string) {
	fileName, err := dialog.File().Title("Save As").Filter("All Files", "*").Save()
	fmt.Println(fileName)
	fmt.Println("Error:", err)

	binaryBody := []byte(stringBody)
	ioutil.WriteFile(fileName, binaryBody, os.ModePerm)
}
