package main

import (
	"flag"

	"github.com/codeskyblue/go-sh"
)

func conversionSearchVersion(version string) string {
	switch version {
	case "2.2":
		return "2.2_r1.1"
	case "2.3":
		return "2.3_r1.10"
	case "2.3.7":
		return "2.3.7_r1.0"
	case "4.0.1":
		return "4.0.1_r1.0"
	case "4.0.3":
		return "4.0.3_r1.0"
	case "4.1.1":
		return "4.1.1_r1.0"
	case "4.1.2":
		return "4.1.2_r1.0"
	case "4.2.0":
		return "4.2.0_r1.0"
	case "4.3.0":
		return "4.3.0_r3.1"
	case "4.3.1":
		return "4.3.1_r1.0"
	case "4.4.0":
		return "4.4.0_r1.0"
	case "4.4.1":
		return "4.4.1_r1.0"
	case "4.4.2":
		return "4.4.2_r1.0"
	case "4.4.3":
		return "4.4.3_r1.0"
	case "4.4.4":
		return "4.4.4_r1.0"
	case "4.4w":
		return "4.4w_r1.0"
	case "5.0.0":
		return "5.0.0_r2.0"
	case "5.0.1":
		return "5.0.1_r1.0"
	case "5.1.0":
		return "5.1.0_r1.0"
	case "5.1.1":
		return "5.1.1_r9.0"
	case "6":
	case "6.0":
		return "6.0.0_r1.0"
	case "7":
	case "7.0":
		return "7.0.0_r1.0"
	}
	return "7.0.0_r1.0"
}

func main() {
	var fullSearch = flag.String("s", "", "Full Search")
	var filePath = flag.String("f", "", "File Path")
	var definition = flag.String("d", "", "Definition")
	var symbol = flag.String("sy", "", "Symbol")
	var version = flag.String("v", "7.0.0_r1.0", "Version")
	flag.Parse()
	sh.Command("open", "http://tools.oesf.biz/android-" + conversionSearchVersion(*version) + "/search?q=" + *fullSearch + "&defs=" + *definition + "&refs=" + *symbol + "&path=" + *filePath + "&hist=").Run()
}
