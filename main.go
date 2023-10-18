package main

import (
	"flag"
	"log"
	"github.com/RedMapleTech/machodump/entitlements"
	"github.com/RedMapleTech/machodump/helpers"
	"os"

	"github.com/blacktop/go-macho"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "i", "", "input file")
	flag.Parse()

	// check arg
	if inputFile == "" {
		flag.Usage()
		log.Fatalf("Must supply input file")
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Fatal: input does not exist: %s", inputFile)
	}

	// process file details
	log.Printf("Parsing file %q.", inputFile)

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("Error opening file %s: %q\n", inputFile, err.Error())
	}

	machoFile, err := macho.NewFile(file)

	if err != nil {
		log.Printf("Error parsing file as Macho-O: %q. Exiting.", err.Error())
		os.Exit(-1)
	}

	// get code signature
	cd := machoFile.CodeSignature()

	if cd == nil {
		// print file details
		helpers.PrintFileDetails(machoFile)
		helpers.PrintLibs(machoFile.ImportedLibraries())
		helpers.PrintLoads(machoFile.Loads)

		log.Fatalln("No code signing section in binary, exiting")
	}

	// get array of entitlements
	ents, err := entitlements.GetEntsFromXMLString(cd.Entitlements)

	if err != nil {
		log.Printf("Errored when trying to extract entitlements: %s", err.Error())
	}

	// print file details
	helpers.PrintFileDetails(machoFile)
	helpers.PrintLibs(machoFile.ImportedLibraries())
	helpers.PrintLoads(machoFile.Loads)

	// print the details
	helpers.PrintCDs(cd.CodeDirectories)

	// parse the CMS sig, if it's there
	if len(cd.CMSSignature) > 0 {
		helpers.ParseCMSSig(cd.CMSSignature)
	}

	helpers.PrintRequirements(cd.Requirements)

	helpers.PrintEnts(ents)

	log.Printf("Fin.")
}
