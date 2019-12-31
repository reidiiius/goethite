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

func main() {

	var (
		diadem, phyla, cargo string
		pegbox, menu         []string
		tuning               func(string) []string
		epoch                int64
	)

	sdr.ScaleList["z0"] = strings.Repeat("____ ", 12)

	menu = sdr.Signos()

	if !sort.StringsAreSorted(menu) {
		sort.Strings(menu)
	}

	if len(os.Args) < 2 {
		tabular(menu)
	}

	signat, _ := regexp.MatchString(`^[jkn][0-7]+([jknxy][1-7]+)*$`, os.Args[1])

	if !signat {
		diadem = os.Args[1]
		os.Args = os.Args[1:]

		if len(os.Args) < 2 {
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
		fmt.Println("\n\t" + os.Args[0] + " ?")
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

	cipher := map[string]string{
		"Hg": "v9",
		"Pu": "zE",
		"Sn": "t7",
		"Mn": "p3",
		"Ur": "xC",
		"Cu": "r5",
		"Pb": "wA",
		"Au": "u8",
		"Np": "yD",
		"Ag": "s6",
		"Ti": "o2",
		"Fe": "q4",
	}

	result := phrase

	for face, veil := range cipher {
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
