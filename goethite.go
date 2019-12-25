package main

import (
	"fmt"
	"github.com/reidiiius/goethite/scordatura"
)

func main() {

	aragonite := [42]string{
		"n0",
		"k6",
		"j17",
		"k6x5",
		"j17y2",
		"j3",
		"j34k6",
		"j17k2",
		"n26y5",
		"k26x5",
		"j6",
		"j36",
		"k56",
		"j136y7",
		"k56x4",
		"n167x4",
		"j3k5x4",
		"j167y2",
		"j2",
		"j236",
		"j26",
		"j23",
		"j23k6",
		"j2y3",
		"j2k6",
		"j26y3",
		"j2k56",
		"j246y3",
		"j2k56x4",
		"k157x6",
		"j26y34",
		"j2k6x5",
		"j2k6y3",
		"k1j6",
		"n345",
		"j3k6",
		"n45y2",
		"j3k56x4",
		"k2j6",
		"n5y2",
		"k26",
		"k256",
	}

	p := fmt.Println
	f := scordatura.EADGBE // Tuning

	for i := 0; i < len(aragonite); i++ {
		a := f(aragonite[i])
		p("")
		for j := 0; j < len(a); j++ {
			p("\t" + a[j])
		}
	}

	p("")

}
