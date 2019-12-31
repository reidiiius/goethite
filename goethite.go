package main

import (
	"fmt"
	sdr "github.com/reidiiius/goethite/scordatura"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {

	var (
		diadem, phyla, cargo string
		menu, pegbox         []string
		tuning               func(string) []string
		epoch                int64
	)

	sdr.ScaleList["z0"] = strings.Repeat("____ ", 12)

	menu = sdr.Signos()

	if !sort.StringsAreSorted(menu) {
		sort.Strings(menu)
	}

	if len(os.Args) < 2 {
		for i := 0; i < len(menu); i++ {
			if i%7 == 0 {
				fmt.Println("")
			}
			fmt.Print("\t", menu[i])
		}

		fmt.Print("\n\n")
		os.Exit(0)
	}

	diadem = "eadgbe"

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
		diadem = "unison"
		tuning = sdr.Unison
	}

	epoch = time.Now().UnixNano()
	phyla = fmt.Sprintf("-%s-v%d", diadem, epoch)

	for i := 1; i < len(os.Args); i++ {
		cargo = strings.ToLower(os.Args[i])

		if sdr.Absent(cargo) {
			fmt.Println("\n\t" + cargo + " ?")
			continue
		}

		pegbox = tuning(cargo)

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
	metals := []string{
		"Hg", "Pu", "Sn", "Mn", "Ur", "Cu",
		"Pb", "Au", "Np", "Ag", "Ti", "Fe",
	}

	cipher := []string{
		"v9", "zE", "t7", "p3", "xC", "r5",
		"wA", "u8", "yD", "s6", "o2", "q4",
	}

	result := phrase

	for windex, element := range metals {
		result = strings.ReplaceAll(result, element, cipher[windex])
	}

	return result
}
