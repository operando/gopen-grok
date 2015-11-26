package main

import (
	"flag"

	"github.com/codeskyblue/go-sh"
)

func main() {
	var fullSerch = flag.String("s", "", "Full Search")
	var filePath = flag.String("f", "", "File Path")
	var definition = flag.String("d", "", "Definition")
	var symbol = flag.String("sy", "", "Symbol")
	var version = flag.String("v", "6.0.0_r1.0", "Version")
	flag.Parse()
	sh.Command("open", "http://tools.oesf.biz/android-"+*version+"/search?q="+*fullSerch+"&defs="+*definition+"&refs="+*symbol+"&path="+*filePath+"&hist=").Run()
}
