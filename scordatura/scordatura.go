// Package scordatura provides structured data and routines to model
// tonal inversion functions of musical pitch collections by manipulating
// text to display semigraphic fingerboard diagrams of stringed intruments.
package scordatura

var ScaleList = map[string]string{
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

var PitchFork = map[string]byte{
	"Bk": 0,
	"Cn": 0,

	"Ck": 5,
	"Dj": 5,

	"Dn": 10,

	"Dk": 15,
	"Ej": 15,

	"En": 20,
	"Fj": 20,

	"Ek": 25,
	"Fn": 25,

	"Fk": 30,
	"Gj": 30,

	"Gn": 35,

	"Gk": 40,
	"Aj": 40,

	"An": 45,

	"Ak": 50,
	"Bj": 50,

	"Bn": 55,
	"Cj": 55,
}

// Utilities

func Absent(argot string) bool {
	_, ok := ScaleList[argot]

	return !ok
}

func Signos() []string {
	var keystones []string

	for keynotes := range ScaleList {
		keystones = append(keystones, keynotes)
	}
	return keystones
}

func HeadStock(argot, tone string) string {
	spandex := PitchFork[tone]
	temple := ScaleList[argot]
	course := temple[spandex:] + temple[:spandex]
	return course
}

func LatticeWork(argot string, tuned []string) []string {
	var course string
	var board []string

	board = append(board, argot)

	for _, tone := range tuned {
		course = HeadStock(argot, tone)
		board = append(board, course)
	}
	return board
}

// Tunings

func BEADGCF(argot string) []string {
	tuned := []string{"Fn", "Cn", "Gn", "Dn", "An", "En", "Bn"}
	board := LatticeWork(argot, tuned)
	return board
}

func BFBFB(argot string) []string {
	tuned := []string{"Bn", "Fn", "Bn", "Fn", "Bn"}
	board := LatticeWork(argot, tuned)
	return board
}

func CGDAE(argot string) []string {
	tuned := []string{"En", "An", "Dn", "Gn", "Cn"}
	board := LatticeWork(argot, tuned)
	return board
}

func DADGAD(argot string) []string {
	tuned := []string{"Dn", "An", "Gn", "Dn", "An", "Dn"}
	board := LatticeWork(argot, tuned)
	return board
}

func DADGBD(argot string) []string {
	tuned := []string{"Dn", "Bn", "Gn", "Dn", "An", "Dn"}
	board := LatticeWork(argot, tuned)
	return board
}

func EADGBE(argot string) []string {
	tuned := []string{"En", "Bn", "Gn", "Dn", "An", "En"}
	board := LatticeWork(argot, tuned)
	return board
}

func FkBjDn(argot string) []string {
	tuned := []string{"Dn", "Bj", "Fk", "Dn", "Bj", "Fk"}
	board := LatticeWork(argot, tuned)
	return board
}

func Unison(argot string) []string {
	tuned := []string{"Cn"}
	board := LatticeWork(argot, tuned)
	return board
}
