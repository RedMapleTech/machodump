package main

import (
	"flag"
	"log"
	"redmapletech/machodump/entitlements"
	"redmapletech/machodump/helpers"
	"os"

	"github.com/blacktop/go-macho"
)

func main() {
	inputPtr := flag.String("i", "", "input file")
	flag.Parse()

	// check arg
	if *inputPtr == "" {
		flag.Usage()
		panic("Must supply input file")
	}

	input := *inputPtr

	if _, err := os.Stat(input); os.IsNotExist(err) {
		log.Panicf("Fatal: input does not exist: %s", input)
	}

	// process file details
	log.Printf("Parsing file %q.", input)

	file, err := os.Open(input)

	if err != nil {
		panic(err)
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

		log.Printf("No code signing section in binary, exiting")
		os.Exit(1)
	}

	// get array of entitlements
	ents, err := entitlements.GetEntsFromXMLString(cd.Entitlements)

	if err != nil {
		log.Printf("Errored when trying to extract ents: %s", err.Error())
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
