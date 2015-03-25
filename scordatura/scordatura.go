package main

import "fmt"

var clave = map[string]string{
	"j2":      "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu ",
	"j3":      "HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr ",
	"j5":      "PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb ",
	"j6":      "HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp ",
	"k1":      "____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ",
	"k2":      "NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
	"k5":      "UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
	"k6":      "HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi ",
	"n0":      "HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ FePb ",
	"j17":     "____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
	"j23":     "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
	"j25":     "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb ",
	"j26":     "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
	"j36":     "HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp ",
	"j56":     "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb ",
	"k12":     "____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
	"k15":     "____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
	"k25":     "NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
	"k26":     "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu ",
	"k34":     "PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb ",
	"k56":     "HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp ",
	"j236":    "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
	"j256":    "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb ",
	"j2k5":    "FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb ",
	"j2k6":    "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu ",
	"j2y3":    "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
	"j3k5":    "NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb ",
	"j3k6":    "HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe ",
	"j5y6":    "PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb ",
	"k125":    "____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
	"k1j5":    "____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ",
	"k1j6":    "____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
	"k256":    "HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu ",
	"k2j5":    "NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
	"k2j6":    "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu ",
	"k2x1":    "____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
	"k6x5":    "HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi ",
	"n167":    "NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb ",
	"n345":    "____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ",
	"n5y2":    "HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu ",
	"n6x2":    "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ ",
	"j17k2":   "____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ",
	"j17y2":   "HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
	"j23k6":   "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
	"j25y6":   "TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb ",
	"j26y3":   "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
	"j2k34":   "TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb ",
	"j2k56":   "HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu ",
	"j34k6":   "HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn ",
	"j56y7":   "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ ",
	"k12j5":   "____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
	"k17j5":   "TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ ",
	"k25x1":   "____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
	"k26x5":   "HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
	"k2j56":   "NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb ",
	"k34x2":   "PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb ",
	"k56x4":   "HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp ",
	"n25x6":   "TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb ",
	"n26y5":   "____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
	"n45y2":   "HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe ",
	"n67x2":   "____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ",
	"j136y7":  "____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
	"j167y2":  "HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp ",
	"j246y3":  "HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu ",
	"j26y34":  "HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu ",
	"j2k6x5":  "HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu ",
	"j2k6y3":  "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
	"j346y5":  "NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb ",
	"j3k5x4":  "HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp ",
	"k135x4":  "____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ",
	"k157x6":  "HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu ",
	"k1j6y7":  "____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ",
	"k257x1":  "NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
	"k25x17":  "____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
	"k2j5x1":  "____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
	"k2j5y6":  "NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb ",
	"k345x2":  "UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb ",
	"n167x4":  "HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
	"n345y7":  "____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ",
	"j2k56x4": "HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu ",
	"j3k56x4": "HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe ",
	"k1j56y7": "____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ",
	"k2j56y7": "NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ "}

func fn(tuner string) string {
	return (clave[tuner][25:60] + clave[tuner][0:25])
}
func cn(tuner string) string {
	return (clave[tuner])
}
func gn(tuner string) string {
	return (clave[tuner][35:60] + clave[tuner][0:35])
}
func dn(tuner string) string {
	return (clave[tuner][10:60] + clave[tuner][0:10])
}
func an(tuner string) string {
	return (clave[tuner][45:60] + clave[tuner][0:45])
}
func en(tuner string) string {
	return (clave[tuner][20:60] + clave[tuner][0:20])
}
func bn(tuner string) string {
	return (clave[tuner][55:60] + clave[tuner][0:55])
}

func main() {
	p := fmt.Println

	p("\nn0")
	p(fn("n0"))
	p(cn("n0"))
	p(gn("n0"))
	p(dn("n0"))
	p(an("n0"))
	p(en("n0"))
	p(bn("n0"))

	p("\nk6")
	p(fn("k6"))
	p(cn("k6"))
	p(gn("k6"))
	p(dn("k6"))
	p(an("k6"))
	p(en("k6"))
	p(bn("k6"))

	p("\nj17")
	p(fn("j17"))
	p(cn("j17"))
	p(gn("j17"))
	p(dn("j17"))
	p(an("j17"))
	p(en("j17"))
	p(bn("j17"))

	p("\nk6x5")
	p(fn("k6x5"))
	p(cn("k6x5"))
	p(gn("k6x5"))
	p(dn("k6x5"))
	p(an("k6x5"))
	p(en("k6x5"))
	p(bn("k6x5"))

	p("\nj17y2")
	p(fn("j17y2"))
	p(cn("j17y2"))
	p(gn("j17y2"))
	p(dn("j17y2"))
	p(an("j17y2"))
	p(en("j17y2"))
	p(bn("j17y2"))

	p("\nj3")
	p(fn("j3"))
	p(cn("j3"))
	p(gn("j3"))
	p(dn("j3"))
	p(an("j3"))
	p(en("j3"))
	p(bn("j3"))

	p("\nj34k6")
	p(fn("j34k6"))
	p(cn("j34k6"))
	p(gn("j34k6"))
	p(dn("j34k6"))
	p(an("j34k6"))
	p(en("j34k6"))
	p(bn("j34k6"))

	p("\nj17k2")
	p(fn("j17k2"))
	p(cn("j17k2"))
	p(gn("j17k2"))
	p(dn("j17k2"))
	p(an("j17k2"))
	p(en("j17k2"))
	p(bn("j17k2"))

	p("\nn26y5")
	p(fn("n26y5"))
	p(cn("n26y5"))
	p(gn("n26y5"))
	p(dn("n26y5"))
	p(an("n26y5"))
	p(en("n26y5"))
	p(bn("n26y5"))

	p("\nk26x5")
	p(fn("k26x5"))
	p(cn("k26x5"))
	p(gn("k26x5"))
	p(dn("k26x5"))
	p(an("k26x5"))
	p(en("k26x5"))
	p(bn("k26x5"))

	p("\nj6")
	p(fn("j6"))
	p(cn("j6"))
	p(gn("j6"))
	p(dn("j6"))
	p(an("j6"))
	p(en("j6"))
	p(bn("j6"))

	p("\nj36")
	p(fn("j36"))
	p(cn("j36"))
	p(gn("j36"))
	p(dn("j36"))
	p(an("j36"))
	p(en("j36"))
	p(bn("j36"))

	p("\nk56")
	p(fn("k56"))
	p(cn("k56"))
	p(gn("k56"))
	p(dn("k56"))
	p(an("k56"))
	p(en("k56"))
	p(bn("k56"))

	p("\nj136y7")
	p(fn("j136y7"))
	p(cn("j136y7"))
	p(gn("j136y7"))
	p(dn("j136y7"))
	p(an("j136y7"))
	p(en("j136y7"))
	p(bn("j136y7"))

	p("\nk56x4")
	p(fn("k56x4"))
	p(cn("k56x4"))
	p(gn("k56x4"))
	p(dn("k56x4"))
	p(an("k56x4"))
	p(en("k56x4"))
	p(bn("k56x4"))

	p("\nn167x4")
	p(fn("n167x4"))
	p(cn("n167x4"))
	p(gn("n167x4"))
	p(dn("n167x4"))
	p(an("n167x4"))
	p(en("n167x4"))
	p(bn("n167x4"))

	p("\nj3k5x4")
	p(fn("j3k5x4"))
	p(cn("j3k5x4"))
	p(gn("j3k5x4"))
	p(dn("j3k5x4"))
	p(an("j3k5x4"))
	p(en("j3k5x4"))
	p(bn("j3k5x4"))

	p("\nj167y2")
	p(fn("j167y2"))
	p(cn("j167y2"))
	p(gn("j167y2"))
	p(dn("j167y2"))
	p(an("j167y2"))
	p(en("j167y2"))
	p(bn("j167y2"))

	p("\nj2")
	p(fn("j2"))
	p(cn("j2"))
	p(gn("j2"))
	p(dn("j2"))
	p(an("j2"))
	p(en("j2"))
	p(bn("j2"))

	p("\nj236")
	p(fn("j236"))
	p(cn("j236"))
	p(gn("j236"))
	p(dn("j236"))
	p(an("j236"))
	p(en("j236"))
	p(bn("j236"))

	p("\nj26")
	p(fn("j26"))
	p(cn("j26"))
	p(gn("j26"))
	p(dn("j26"))
	p(an("j26"))
	p(en("j26"))
	p(bn("j26"))

	p("\nj23")
	p(fn("j23"))
	p(cn("j23"))
	p(gn("j23"))
	p(dn("j23"))
	p(an("j23"))
	p(en("j23"))
	p(bn("j23"))

	p("\nj23k6")
	p(fn("j23k6"))
	p(cn("j23k6"))
	p(gn("j23k6"))
	p(dn("j23k6"))
	p(an("j23k6"))
	p(en("j23k6"))
	p(bn("j23k6"))

	p("\nj2y3")
	p(fn("j2y3"))
	p(cn("j2y3"))
	p(gn("j2y3"))
	p(dn("j2y3"))
	p(an("j2y3"))
	p(en("j2y3"))
	p(bn("j2y3"))

	p("\nj2k6")
	p(fn("j2k6"))
	p(cn("j2k6"))
	p(gn("j2k6"))
	p(dn("j2k6"))
	p(an("j2k6"))
	p(en("j2k6"))
	p(bn("j2k6"))

	p("\nj26y3")
	p(fn("j26y3"))
	p(cn("j26y3"))
	p(gn("j26y3"))
	p(dn("j26y3"))
	p(an("j26y3"))
	p(en("j26y3"))
	p(bn("j26y3"))

	p("\nj2k56")
	p(fn("j2k56"))
	p(cn("j2k56"))
	p(gn("j2k56"))
	p(dn("j2k56"))
	p(an("j2k56"))
	p(en("j2k56"))
	p(bn("j2k56"))

	p("\nj246y3")
	p(fn("j246y3"))
	p(cn("j246y3"))
	p(gn("j246y3"))
	p(dn("j246y3"))
	p(an("j246y3"))
	p(en("j246y3"))
	p(bn("j246y3"))

	p("\nj2k56x4")
	p(fn("j2k56x4"))
	p(cn("j2k56x4"))
	p(gn("j2k56x4"))
	p(dn("j2k56x4"))
	p(an("j2k56x4"))
	p(en("j2k56x4"))
	p(bn("j2k56x4"))

	p("\nk157x6")
	p(fn("k157x6"))
	p(cn("k157x6"))
	p(gn("k157x6"))
	p(dn("k157x6"))
	p(an("k157x6"))
	p(en("k157x6"))
	p(bn("k157x6"))

	p("\nj26y34")
	p(fn("j26y34"))
	p(cn("j26y34"))
	p(gn("j26y34"))
	p(dn("j26y34"))
	p(an("j26y34"))
	p(en("j26y34"))
	p(bn("j26y34"))

	p("\nj2k6x5")
	p(fn("j2k6x5"))
	p(cn("j2k6x5"))
	p(gn("j2k6x5"))
	p(dn("j2k6x5"))
	p(an("j2k6x5"))
	p(en("j2k6x5"))
	p(bn("j2k6x5"))

	p("\nj2k6y3")
	p(fn("j2k6y3"))
	p(cn("j2k6y3"))
	p(gn("j2k6y3"))
	p(dn("j2k6y3"))
	p(an("j2k6y3"))
	p(en("j2k6y3"))
	p(bn("j2k6y3"))

	p("\nk1j6")
	p(fn("k1j6"))
	p(cn("k1j6"))
	p(gn("k1j6"))
	p(dn("k1j6"))
	p(an("k1j6"))
	p(en("k1j6"))
	p(bn("k1j6"))

	p("\nn345")
	p(fn("n345"))
	p(cn("n345"))
	p(gn("n345"))
	p(dn("n345"))
	p(an("n345"))
	p(en("n345"))
	p(bn("n345"))

	p("\nj3k6")
	p(fn("j3k6"))
	p(cn("j3k6"))
	p(gn("j3k6"))
	p(dn("j3k6"))
	p(an("j3k6"))
	p(en("j3k6"))
	p(bn("j3k6"))

	p("\nn45y2")
	p(fn("n45y2"))
	p(cn("n45y2"))
	p(gn("n45y2"))
	p(dn("n45y2"))
	p(an("n45y2"))
	p(en("n45y2"))
	p(bn("n45y2"))

	p("\nj3k56x4")
	p(fn("j3k56x4"))
	p(cn("j3k56x4"))
	p(gn("j3k56x4"))
	p(dn("j3k56x4"))
	p(an("j3k56x4"))
	p(en("j3k56x4"))
	p(bn("j3k56x4"))

	p("\nk2j6")
	p(fn("k2j6"))
	p(cn("k2j6"))
	p(gn("k2j6"))
	p(dn("k2j6"))
	p(an("k2j6"))
	p(en("k2j6"))
	p(bn("k2j6"))

	p("\nn5y2")
	p(fn("n5y2"))
	p(cn("n5y2"))
	p(gn("n5y2"))
	p(dn("n5y2"))
	p(an("n5y2"))
	p(en("n5y2"))
	p(bn("n5y2"))

	p("\nk26")
	p(fn("k26"))
	p(cn("k26"))
	p(gn("k26"))
	p(dn("k26"))
	p(an("k26"))
	p(en("k26"))
	p(bn("k26"))

	p("\nk256")
	p(fn("k256"))
	p(cn("k256"))
	p(gn("k256"))
	p(dn("k256"))
	p(an("k256"))
	p(en("k256"))
	p(bn("k256"))

	p("")
}
