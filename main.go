package main

import (
	"aqwari.net/xml/xsdgen"
	"log"
	"os"
)

func main() {
	var cfg xsdgen.Config
	cfg.Option(
		xsdgen.PackageName("main"),
		xsdgen.LogOutput(log.New(os.Stderr, "", 0)),
		xsdgen.LogLevel(2),
	)
	err := cfg.GenCLI("pacs.008.001.09.xsd")
	if err != nil {
		log.Fatal(err)
	}

}
