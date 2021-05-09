package main

import (
	core "github.com/akakou-hobby/uree-core-packages"
	uree_package "github.com/akakou-hobby/uree-package"
)

func navber_packages() []uree_package.UreeNavberPackage {
	return []uree_package.UreeNavberPackage{core.CommpileCPackage{}}
}

func side_pallet_packages() []uree_package.UreeLeftPackage {
	return []uree_package.UreeLeftPackage{core.FileSidePallet{}}
}
