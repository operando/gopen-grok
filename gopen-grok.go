package main

import (
	"flag"

	"github.com/codeskyblue/go-sh"
)

func conversionSearchVersion(version string) string {
	switch version {
	case "6":
	case "6.0":
		return "6.0.0_r1.0"
	}
	return "7.0.0_r1.0"
}

func main() {
	var fullSerch = flag.String("s", "", "Full Search")
	var filePath = flag.String("f", "", "File Path")
	var definition = flag.String("d", "", "Definition")
	var symbol = flag.String("sy", "", "Symbol")
	var version = flag.String("v", "7.0.0_r1.0", "Version")
	flag.Parse()
	sh.Command("open", "http://tools.oesf.biz/android-"+conversionSearchVersion(*version)+"/search?q="+*fullSerch+"&defs="+*definition+"&refs="+*symbol+"&path="+*filePath+"&hist=").Run()
}
