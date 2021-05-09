package main

import (
	"fmt"
	"strings"

	core "github.com/akakou-hobby/uree-core-packages"
	uree_package "github.com/akakou-hobby/uree-package"
)

type UreeOnChangeEventPackage interface {
	Run(uree_package.Request) uree_package.Response
	GetName() string
}

type IntHighter struct {
}

func (pkg IntHighter) Run(req uree_package.Request) uree_package.Response {
	old_body := req.Body
	new_body := strings.Replace(old_body, "int", "<p style='color=#79aeff'>int</p>", -1)
	new_body = strings.Replace(new_body, "<p style='color=#79aeff'><p style='color=#79aeff'>int</p></p>", "<p style='color=#79aeff'>int</p>", -1)

	fmt.Println(new_body)

	return uree_package.Response{new_body}
}

func (pkg IntHighter) GetName() string {
	return ""
}

func navber_packages() []uree_package.UreeNavberPackage {
	return []uree_package.UreeNavberPackage{core.CommpileCPackage{}}
}

func side_pallet_packages() []uree_package.UreeLeftPackage {
	return []uree_package.UreeLeftPackage{core.FileSidePallet{}}
}

func on_change_event_packages() []UreeOnChangeEventPackage {
	return []UreeOnChangeEventPackage{IntHighter{}}
}
