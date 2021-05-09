//go:generate fileb0x b0x.yml
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ImVexed/muon"
	uree_package "github.com/akakou-hobby/uree-package"
	"github.com/sqweek/dialog"
)

var filePath string
var m *muon.Window
var navber_pkgs []uree_package.UreeNavberPackage

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

	m = muon.New(cfg, fileHandler)

	m.Bind("openFile", openFile)
	m.Bind("saveFile", saveFile)
	m.Bind("saveFileAsNew", saveFileAsNew)
	m.Bind("loadNavbar", loadNavbar)

	navber_pkgs = navber_packages()

	loadNavbar()

	if err := m.Start(); err != nil {
		panic(err)
	}
}

func openFile() string {
	path, err := dialog.File().Title("Open").Filter("All Files", "*").Load()

	if err != nil {
		dialog.Message("%s", err).Title("Error").Info()
		return ""
	}

	filePath = path

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	stringBody := string(bytes)
	stringBody = strings.Replace(stringBody, " ", " ", -1)

	return stringBody
}

func saveFile(stringBody string) {
	fmt.Println("%d saved!", filePath)
	fmt.Println(stringBody)

	stringBody = strings.Replace(stringBody, " ", " ", -1)
	binaryBody := []byte(stringBody)
	ioutil.WriteFile(filePath, binaryBody, os.ModePerm)
}

func saveFileAsNew(stringBody string) {
	path, err := dialog.File().Title("Save As").Filter("All Files", "*").Save()

	fmt.Println(path)
	fmt.Println("Error:", err)

	if err != nil {
		dialog.Message("%s", err).Title("Error").Info()
		return
	}

	filePath = path
	saveFile(stringBody)
}

func loadNavbar() string {
	nabver := ""
	fmt.Println(len(navber_pkgs))

	for i := 0; i < len(navber_pkgs); i++ {
		pkg := navber_pkgs[i]
		name := pkg.GetName()
		optional := pkg.SetUpOptional()

		if optional == "" {
			optional = "null"
		}

		func_name := fmt.Sprintf("navbar_%d", i)

		f := func(body string, optional string) string {
			req := uree_package.Request{
				Path:     filePath,
				Body:     body,
				Optional: optional,
			}

			resp := pkg.Run(req)
			return resp.Body
		}

		m.Bind(func_name, f)
		fmt.Println(optional)

		nabver += fmt.Sprintf("<li><a onclick=\"document.body.innerHTML=%s(document.body.innerHTML, '${%s}')\">%s</a></li>", func_name, optional, name)
	}

	fmt.Println(nabver)

	return nabver
}

func loadLeft() string {
	nabver := ""
	fmt.Println(len(navber_pkgs))

	for i := 0; i < len(navber_pkgs); i++ {
		pkg := navber_pkgs[i]
		name := pkg.GetName()
		optional := pkg.SetUpOptional()
	}

	fmt.Println(nabver)

	return nabver
}
