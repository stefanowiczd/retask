package main

import (
	"github.com/stefanowiczd/retask/internal/application"
	"github.com/stefanowiczd/retask/internal/interface/rest"
)

func main() {

	pkgMgr := application.NewPackageManager()

	svr := rest.NewServer(rest.DefaultConfig(), pkgMgr)

	svr.Start()
}
