package main

import (
	"fmt"
	"os/exec"

	uree_package "github.com/akakou-hobby/uree-package"
	"github.com/sqweek/dialog"
)

type CommpileCPackage struct{}

func (pkg CommpileCPackage) Run(req uree_package.Request) uree_package.Response {
	resp := uree_package.Response{
		req.Body,
	}

	outputPath := fmt.Sprintf("%s.out", req.Path)

	fmt.Printf("gcc %s -o %s", req.Path, outputPath)

	out, err := exec.Command("gcc", req.Path, "-o", outputPath).CombinedOutput()
	if err != nil {
		dialog.Message("Error", string(out)).Title("Compile error").Info()
		return resp
	}

	err = exec.Command("chmod", "+x", outputPath).Run()
	if err != nil {
		dialog.Message("Error", err).Title("chmod Error").Info()
		return resp
	}

	out, err = exec.Command(outputPath).Output()

	dialog.Message("Output", string(out)).Title("Output").Info()

	fmt.Println(req.Body)
	fmt.Println("hello")

	return resp
}

func (pkg CommpileCPackage) SetUpOptional() string {
	return ""
}

func (pkg CommpileCPackage) GetName() string {
	return "Run C"
}

// type PasteNavberPackage struct{}

// func (pkg CutNavbarPackage) Run(req uree_package.Request) uree_package.Response {
// 	clipboard = req.Optional

// 	return uree_package.Response{
// 		req.Body,
// 	}
// }

func cut() {

}

func navber_packages() []NavberUreePackage {
	return []NavberUreePackage{CommpileCPackage{}}
}
