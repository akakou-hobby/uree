package main

import (
	core "github.com/akakou-hobby/uree-core-packages"
	uree_package "github.com/akakou-hobby/uree-package"
)

func navber_packages() []uree_package.UreeNavberPackage {
	return []uree_package.UreeNavberPackage{core.CommpileCPackage{}}
}

type FileSidePallet struct {
}

func (pkg FileSidePallet) Run(req uree_package.Request) uree_package.Response {
	return uree_package.Response{""}
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

func side_pallet_packages() []uree_package.UreeLeftPackage {
	return []uree_package.UreeLeftPackage{FileSidePallet{}}
}
