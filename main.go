package main

import (
	"log"

	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
)

func main() {
	pkg, _ := apk.OpenFile("app.apk")
	defer pkg.Close()
	icon, _ := pkg.Icon(nil)
	pkgName := pkg.PackageName()
	resConfigEN := &androidbinary.ResTableConfig{
		Language: [2]uint8{uint8('z'), uint8('h')},
	}
	// appLabel :=
	appLabel, _ := pkg.Label(resConfigEN)

	log.Printf("  icon %q...", icon)
	log.Printf("  pkgName %q...", pkgName)
	log.Printf("  appLabel %q...", appLabel)
}
