package main

import (
	"fmt"
	"os/exec"
	"strings"

	core "github.com/akakou-hobby/uree-core-packages"
	uree_package "github.com/akakou-hobby/uree-package"
)

type FileSidePallet struct {
}

func (pkg FileSidePallet) Run(req uree_package.Request) uree_package.Response {
	divided := strings.Split(filePath, "/")
	divided = divided[:len(divided)-1]

	parent := strings.Join(divided, "/")
	fmt.Print(parent)

	out, _ := exec.Command("tree", parent).Output()

	filtered_out := strings.Replace(string(out[:10000]), " ", "&ensp;", -1)
	filtered_out = strings.Replace(filtered_out, "\n", "<br>", -1)

	return uree_package.Response{filtered_out}
}

func (pkg FileSidePallet) SetUpOptional() string {
	return ""
}

func (pkg FileSidePallet) GetName() string {
	return ""
}

func (pkg FileSidePallet) GetIconPath() string {
	return "./img/file.png"
}

func navber_packages() []uree_package.UreeNavberPackage {
	return []uree_package.UreeNavberPackage{core.CommpileCPackage{}}
}

func side_pallet_packages() []uree_package.UreeLeftPackage {
	return []uree_package.UreeLeftPackage{FileSidePallet{}}
}
