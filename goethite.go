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
		diadem, phyla, barge string
		pegbox, cargo, menu  []string
		epoch                int64
		items                int
	)

	qp := &sdr.Schematic

	qp.Septets["z0"] = strings.Repeat("____ ", 12)

	menu = qp.Signos()

	if !sort.StringsAreSorted(menu) {
		sort.Strings(menu)
	}

	cargo = make([]string, len(os.Args))
	items = copy(cargo, os.Args[0:])

	if items < 2 {
		tabular(menu)
	}

	barge = strings.ToLower(cargo[1])

	pegbox = []string{
		"beadgcf",
		"bfbfb",
		"cgdae",
		"dadgad",
		"dadgbd",
		"eadgbe",
		"fkbjdn",
		"unison",
	}

	signat, _ := regexp.MatchString(`^[jknz]\d+([jknxy]\d+)*$`, barge)

	if !signat {
		diadem = barge
		cargo = cargo[1:]

		if len(cargo) < 2 {
			tabular(menu)
		}
	} else {
		diadem = pegbox[5]
	}

	switch diadem {
	case pegbox[0]:
	case pegbox[1]:
	case pegbox[2]:
	case pegbox[3]:
	case pegbox[4]:
	case pegbox[5]:
	case pegbox[6]:
	case pegbox[7]:
	default:
		fmt.Printf("\n\t%s ?\n\n", cargo[0])

		for _, v := range pegbox {
			fmt.Printf("\t\t%s\n", v)
		}
		diadem = pegbox[7]
	}

	epoch = time.Now().UnixNano()
	phyla = fmt.Sprintf("-%s-v%d", diadem, epoch)

	for i := 1; i < len(cargo); i++ {
		barge = strings.ToLower(cargo[i])

		if qp.Absent(barge) {
			fmt.Println("\n\t" + barge + " ?")
			continue
		}

		pegbox = qp.Chordophone(barge, diadem)

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

func transcribe(phrase string) string {
	result := phrase
	qp := &sdr.Schematic

	for face, veil := range qp.Decans {
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
