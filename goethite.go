package main

import (
	"fmt"
	"github.com/reidiiius/goethite/scordatura"
	"os"
	"sort"
)

func main() {

	scordatura.ScaleList["z0"] = scordatura.Zilch()

	var menu []string
	menu = scordatura.Signos()

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

		fmt.Println("")
	}

	tuning := scordatura.EADGBE

	for i := 1; i < len(os.Args); i++ {
		if scordatura.Absent(os.Args[i]) {
			fmt.Println("\n\t" + os.Args[i] + " ?")
			continue
		}

		pegbox := tuning(os.Args[i])

		fmt.Println("")
		for j := 0; j < len(pegbox); j++ {
			fmt.Println("\t" + pegbox[j])
		}
	}

	fmt.Println("")
}
