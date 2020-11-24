package helpers

import (
	"fmt"
	"redmapletech/machodump/entitlements"
	"strings"
	"unicode"

	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"
	"github.com/dustin/go-humanize/english"

	ctypes "github.com/blacktop/go-macho/pkg/codesign/types"
)

// PrintFileDetails prints the general file details to console
func PrintFileDetails(m *macho.File) {
	fmt.Printf("File Details:\n"+
		"\tMagic: %s\n"+
		"\tType: %s\n"+
		"\tCPU: %s, %s\n"+
		"\tCommands: %d (Size: %d)\n"+
		"\tFlags: %s\n"+
		"\tUUID: %s\n",
		m.FileHeader.Magic,
		m.FileHeader.Type,
		m.FileHeader.CPU, m.FileHeader.SubCPU.String(m.FileHeader.CPU),
		m.FileHeader.NCommands,
		m.FileHeader.SizeCommands,
		m.FileHeader.Flags.Flags(),
		m.UUID())
}

// PrintLibs prints the loaded libraries to console
func PrintLibs(libs []string) {
	fmt.Printf("File imports %s\n", english.Plural(len(libs), "library:", "libraries:"))

	for i, lib := range libs {
		fmt.Printf("\t%d: %q\n", i, lib)
	}
}

// PrintLoads prints the interesting load commands to console
func PrintLoads(loads []macho.Load) {
	fmt.Printf("File has %s. Interesting %s:\n", english.Plural(len(loads), "load command", "load commands"), english.PluralWord(len(loads), "command", "commands"))

	for i, load := range loads {
		switch load.Command() {
		case types.LC_VERSION_MIN_IPHONEOS:
			fallthrough
		case types.LC_ENCRYPTION_INFO:
			fallthrough
		case types.LC_ENCRYPTION_INFO_64:
			fallthrough
		case types.LC_SOURCE_VERSION:
			fmt.Printf("\tLoad %d (%s): %s\n", i, load.Command(), load.String())
		}
	}
}

// PrintCDs prints the code directory details to console
func PrintCDs(CDs []ctypes.CodeDirectory) {
	fmt.Printf("Binary has %s:\n", english.Plural(len(CDs), "Code Directory", "Code Directories"))

	for i, dir := range CDs {
		fmt.Printf("\tCodeDirectory %d:\n", i)

		fmt.Printf("\t\tIdent: \"%s\"\n", dir.ID)

		if len(dir.TeamID) > 0 && isASCII(dir.TeamID) {
			fmt.Printf("\t\tTeam ID: %q\n", dir.TeamID)
		}

		fmt.Printf("\t\tCD Hash: %s\n", dir.CDHash)
		fmt.Printf("\t\tCode slots: %d\n", len(dir.CodeSlots))
		fmt.Printf("\t\tSpecial slots: %d\n", len(dir.SpecialSlots))

		for _, slot := range dir.SpecialSlots {
			fmt.Printf("\t\t\t%s\n", slot.Desc)
		}
	}
}

// PrintRequirements prints the requirement sections to console
func PrintRequirements(reqs []ctypes.Requirement) {
	fmt.Printf("Binary has %s:\n", english.Plural(len(reqs), "requirement", "requirements"))

	for i, req := range reqs {
		fmt.Printf("\tRequirement %d (%s): %s\n", i, req.Type, req.Detail)
	}
}

// PrintEnts prints the entitlements to console
func PrintEnts(ents *entitlements.EntsStruct) {

	if ents == nil {
		fmt.Printf("Binary has no entitlements\n")
		return
	}

	entries := false

	// print the boolean entries
	if ents.BooleanValues != nil && len(ents.BooleanValues) > 0 {
		fmt.Printf("Binary has %s:\n", english.Plural(len(ents.BooleanValues), "boolean entitlement", "boolean entitlements"))

		for _, ent := range ents.BooleanValues {
			fmt.Printf("\t%s: %t\n", ent.Name, ent.Value)
		}

		entries = true
	}

	// print the string entries
	if ents.StringValues != nil && len(ents.StringValues) > 0 {
		fmt.Printf("Binary has %s:\n", english.Plural(len(ents.StringValues), "string entitlement", "string entitlements"))

		for i, ent := range ents.StringValues {
			fmt.Printf("\t%d %s: %q\n", i, ent.Name, ent.Value)
		}

		entries = true
	}

	// print the integer entries
	if ents.IntegerValues != nil && len(ents.IntegerValues) > 0 {
		fmt.Printf("Binary has %s:\n", english.Plural(len(ents.IntegerValues), "integer entitlement", "integer entitlements"))

		for i, ent := range ents.IntegerValues {
			fmt.Printf("\t%d %s: %d\n", i, ent.Name, ent.Value)
		}

		entries = true
	}

	// print the string array entries
	if ents.StringArrayValues != nil && len(ents.StringArrayValues) > 0 {
		fmt.Printf("Binary has %s:\n", english.Plural(len(ents.StringArrayValues), "string array entitlement", "string array entitlements"))

		for i, ent := range ents.StringArrayValues {

			valueList := ""

			for _, str := range ent.Values {
				valueList = valueList + str + ", "
			}

			valueList = strings.TrimSuffix(valueList, ", ")

			fmt.Printf("\t%d %s: [%q]\n", i, ent.Name, valueList)
		}

		entries = true
	}

	if !entries {
		fmt.Printf("Binary has no entitlements\n")
	}
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
