package main

import (
	"fmt"
	sdr "github.com/reidiiius/goethite/scordatura"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Cipher struct {
	arcane, decans map[string]string
}

var schema = Cipher{
	arcane: map[string]string{
		"Hg": "9",
		"Pu": "R",
		"Sn": "7",
		"Mn": "3",
		"Ur": "P",
		"Cu": "5",
		"Pb": "N",
		"__": "-",
		"Au": "8",
		"Np": "Q",
		"Ag": "6",
		"Ti": "2",
		"Fe": "4",
	},
	decans: map[string]string{
		"Hg": "v",
		"Pu": "z",
		"Sn": "t",
		"Mn": "p",
		"Ur": "x",
		"Cu": "r",
		"Pb": "w",
		"__": "_",
		"Au": "u",
		"Np": "y",
		"Ag": "s",
		"Ti": "o",
		"Fe": "q",
	},
}

func main() {

	var (
		diadem, phyla, barge string
		pegbox, cargo, menu  []string
		tuning               func(string) []string
		epoch                int64
		items                int
	)

	sdr.ScaleList["z0"] = strings.Repeat("____ ", 12)

	menu = sdr.Signos()

	if !sort.StringsAreSorted(menu) {
		sort.Strings(menu)
	}

	cargo = make([]string, len(os.Args))
	items = copy(cargo, os.Args[0:])

	if items < 2 {
		tabular(menu)
	}

	barge = strings.ToLower(cargo[1])

	signat, _ := regexp.MatchString(`^[jkn]\d+([jknxy]\d+)*$`, barge)

	if !signat {
		diadem = barge
		cargo = cargo[1:]

		if len(cargo) < 2 {
			tabular(menu)
		}
	} else {
		diadem = "eadgbe"
	}

	switch diadem {
	case "beadgcf":
		tuning = sdr.BEADGCF

	case "bfbfb":
		tuning = sdr.BFBFB

	case "cgdae":
		tuning = sdr.CGDAE

	case "dadgad":
		tuning = sdr.DADGAD

	case "dadgbd":
		tuning = sdr.DADGBD

	case "eadgbe":
		tuning = sdr.EADGBE

	case "fkbjdn":
		tuning = sdr.FkBjDn

	default:
		fmt.Printf("\n\t%s ?\n", cargo[0])
		diadem = "unison"
		tuning = sdr.Unison
	}

	epoch = time.Now().UnixNano()
	phyla = fmt.Sprintf("-%s-v%d", diadem, epoch)

	for i := 1; i < len(cargo); i++ {
		barge = strings.ToLower(cargo[i])

		if sdr.Absent(barge) {
			fmt.Println("\n\t" + barge + " ?")
			continue
		}

		pegbox = tuning(barge)

		fmt.Println("")
		for j := 0; j < len(pegbox); j++ {
			if j == 0 {
				fmt.Println("\t" + pegbox[j] + phyla)
			} else {
				fmt.Println("\t" + transcribe(pegbox[j]))
			}
		}
	}

	fmt.Println("")
}

// Subroutines

func transcribe(phrase string) string {

	result := phrase

	for face, veil := range schema.decans {
		result = strings.ReplaceAll(result, face, veil)
	}

	return result
}

func tabular(menu []string) {
	for i := 0; i < len(menu); i++ {
		if i%7 == 0 {
			fmt.Println("")
		}
		fmt.Print("\t", menu[i])
	}

	fmt.Print("\n\n")
	os.Exit(0)
}
