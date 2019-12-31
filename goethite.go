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
		menu, pegbox  []string
		diadem, phyla string
		tuning        func(string) []string
		epoch         int64
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
		if sdr.Absent(os.Args[i]) {
			fmt.Println("\n\t" + os.Args[i] + " ?")
			continue
		}

		pegbox = tuning(os.Args[i])

		fmt.Println("")
		for j := 0; j < len(pegbox); j++ {
			if j == 0 {
				fmt.Println("\t" + pegbox[j] + phyla)
			} else {
				fmt.Println("\t" + pegbox[j])
			}
		}
	}

	fmt.Println("")
}
