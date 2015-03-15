package main

import "fmt"

var (
	n0 = map[string]string{ // 1426353696131
		"fn": "PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ CuHg ",
		"cn": "HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ FePb ",
		"gn": "AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ CuHg PbFe ____ ",
		"dn": "SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ ",
		"an": "AgAu ____ FePb HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ ",
		"en": "CuHg PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ ",
		"bn": "FePb HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ "}

	k6 = map[string]string{ // 1426353696133
		"fn": "PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"cn": "HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi ",
		"gn": "AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ ",
		"an": "____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ ",
		"bn": "FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe "}

	j17 = map[string]string{ // 1426353696143
		"fn": "PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb ",
		"cn": "____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"gn": "AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ ",
		"an": "AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ ",
		"bn": "FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe "}

	k6x5 = map[string]string{ // 1426353696146
		"fn": "PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"cn": "HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi ",
		"gn": "____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ ",
		"an": "AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ ",
		"en": "CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ ",
		"bn": "FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe "}

	j17y2 = map[string]string{ // 1426353696150
		"fn": "PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb ",
		"cn": "HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"gn": "AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ ",
		"dn": "____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ",
		"an": "AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ ",
		"bn": "FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe "}

	j3 = map[string]string{ // 1426353696153
		"fn": "PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ ",
		"cn": "HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr ",
		"gn": "AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ ",
		"dn": "SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ ",
		"an": "AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ ",
		"en": "____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ",
		"bn": "FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ "}

	j34k6 = map[string]string{ // 1426353696155
		"fn": "____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ",
		"cn": "HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn ",
		"gn": "AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ ",
		"dn": "SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ ",
		"an": "____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ",
		"en": "CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe ",
		"bn": "FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu "}

	j17k2 = map[string]string{ // 1426353696157
		"fn": "PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi ",
		"cn": "____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ",
		"gn": "AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ ",
		"dn": "____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ",
		"an": "AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ ",
		"en": "CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe ",
		"bn": "FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu "}

	n26y5 = map[string]string{ // 1426353696159
		"fn": "PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi ",
		"cn": "____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"gn": "____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ",
		"dn": "SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ ",
		"an": "AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ ",
		"en": "CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe ",
		"bn": "FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu "}

	k26x5 = map[string]string{ // 1426353696161
		"fn": "PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi ",
		"cn": "HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"gn": "____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ",
		"dn": "____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ",
		"an": "AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ ",
		"en": "CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe ",
		"bn": "FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu "}

	j6 = map[string]string{ // 1426353696163
		"fn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"gn": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"dn": "SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ ",
		"an": "____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ",
		"en": "CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ "}

	j36 = map[string]string{ // 1426353696165
		"fn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"cn": "HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"gn": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ ",
		"dn": "SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ ",
		"an": "____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ",
		"bn": "FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ "}

	k56 = map[string]string{ // 1426353696167
		"fn": "PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp ",
		"gn": "____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"dn": "SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ",
		"en": "CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg "}

	j136y7 = map[string]string{ // 1426353696169
		"fn": "PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ ",
		"cn": "____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ ",
		"dn": "SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ ",
		"an": "AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ",
		"bn": "FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ "}

	k56x4 = map[string]string{ // 1426353696171
		"fn": "____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp ",
		"gn": "AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ ",
		"dn": "SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ",
		"en": "CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg "}

	n167x4 = map[string]string{ // 1426353696172
		"fn": "PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ ",
		"cn": "HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ ",
		"dn": "____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ",
		"an": "AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ",
		"bn": "FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ "}

	j3k5x4 = map[string]string{ // 1426353696174
		"fn": "____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"cn": "HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ ",
		"dn": "SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ",
		"an": "AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe ",
		"en": "____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ",
		"bn": "FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ "}

	j167y2 = map[string]string{ // 1426353696176
		"fn": "PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr ",
		"cn": "HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp ",
		"gn": "AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ ",
		"dn": "____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ",
		"en": "CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ ",
		"bn": "FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg "}

	j2k56x4 = map[string]string{ // 1426353696178
		"fn": "____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu ",
		"gn": "AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ ",
		"dn": "____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ",
		"en": "CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn "}

	k157x6 = map[string]string{ // 1426353696180
		"fn": "PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"cn": "HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu ",
		"gn": "____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ",
		"dn": "SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ",
		"en": "CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ "}

	k1j6 = map[string]string{ // 1426353696182
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp ",
		"cn": "____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ ",
		"dn": "SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe ",
		"an": "____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ ",
		"bn": "FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ "}

	n345 = map[string]string{ // 1426353696184
		"fn": "PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp ",
		"cn": "____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ",
		"gn": "AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ",
		"an": "____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ ",
		"bn": "FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn "}

	j2 = map[string]string{ // 1426353696186
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ",
		"en": "CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ "}

	j26 = map[string]string{ // 1426353696188
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ "}

	j236 = map[string]string{ // 1426353696190
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePb HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ",
		"en": "____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ "}

	j23 = map[string]string{ // 1426353696192
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ "}

	j23k6 = map[string]string{ // 1426353696194
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn "}

	j2y3 = map[string]string{ // 1426353696196
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ "}

	j2k6 = map[string]string{ // 1426353696198
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ",
		"en": "CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn "}

	j26y3 = map[string]string{ // 1426353696200
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ",
		"en": "____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ "}

	j2k56 = map[string]string{ // 1426353696202
		"fn": "PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu ",
		"gn": "____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ",
		"en": "CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn "}

	j246y3 = map[string]string{ // 1426353696204
		"fn": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"cn": "HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ ",
		"dn": "SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ",
		"en": "CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ "}

	j26y34 = map[string]string{ // 1426353696206
		"fn": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ",
		"cn": "HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ ",
		"dn": "SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ",
		"en": "____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ",
		"bn": "FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ "}

	j2k6x5 = map[string]string{ // 1426353696208
		"fn": "PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu ",
		"gn": "____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ",
		"an": "AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ ",
		"en": "CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn "}

	j2k6y3 = map[string]string{ // 1426353696209
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn "}

	j3k6 = map[string]string{ // 1426353696211
		"fn": "PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ",
		"cn": "HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe ",
		"gn": "AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ ",
		"dn": "SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ",
		"en": "____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ",
		"bn": "FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg "}

	n45y2 = map[string]string{ // 1426353696213
		"fn": "PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ ",
		"cn": "HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe ",
		"gn": "AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ ",
		"dn": "____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ",
		"en": "____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ",
		"bn": "FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg "}

	j3k56x4 = map[string]string{ // 1426353696215
		"fn": "____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ",
		"cn": "HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe ",
		"gn": "AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ ",
		"dn": "SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ",
		"en": "____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ",
		"bn": "FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg "}

	k2j6 = map[string]string{ // 1426353696217
		"fn": "PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu ",
		"gn": "AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ",
		"an": "____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ",
		"en": "CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ "}

	n5y2 = map[string]string{ // 1426353696219
		"fn": "____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu ",
		"gn": "AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ ",
		"dn": "____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ",
		"en": "CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb "}

	k26 = map[string]string{ // 1426353696221
		"fn": "PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu ",
		"gn": "AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ",
		"en": "CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb "}

	k256 = map[string]string{ // 1426353696223
		"fn": "PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu ",
		"gn": "____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ",
		"en": "CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb "}

	j5 = map[string]string{ // 1426353696133
		"fn": "TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg ",
		"cn": "PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb ",
		"gn": "____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"dn": "AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ ",
		"en": "AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ "}

	k34 = map[string]string{ // 1426353696143
		"fn": "TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ ",
		"cn": "PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"gn": "HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi ",
		"dn": "AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ ",
		"en": "____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ "}

	j5y6 = map[string]string{ // 1426353696146
		"fn": "TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg ",
		"cn": "PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb ",
		"gn": "HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"dn": "AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ ",
		"an": "____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ",
		"en": "AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ "}

	k34x2 = map[string]string{ // 1426353696150
		"fn": "TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg ",
		"cn": "PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"gn": "HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi ",
		"dn": "____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ ",
		"en": "AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ ",
		"bn": "CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ "}

	k1 = map[string]string{ // 1426353696153
		"fn": "UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg ",
		"cn": "____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ",
		"gn": "PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ ",
		"dn": "HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr ",
		"an": "AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ ",
		"en": "SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ ",
		"bn": "AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ "}

	k17j5 = map[string]string{ // 1426353696155
		"fn": "MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg ",
		"cn": "TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ ",
		"gn": "____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ",
		"dn": "HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn ",
		"an": "AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ ",
		"en": "SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ ",
		"bn": "____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ "}

	j2k34 = map[string]string{ // 1426353696157
		"fn": "MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ ",
		"cn": "TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb ",
		"gn": "PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi ",
		"dn": "____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ",
		"an": "AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ ",
		"en": "____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ",
		"bn": "AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ "}

	n25x6 = map[string]string{ // 1426353696159
		"fn": "MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ ",
		"cn": "TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb ",
		"gn": "PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi ",
		"dn": "HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"an": "____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ",
		"en": "____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ",
		"bn": "AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ "}

	j25y6 = map[string]string{ // 1426353696161
		"fn": "MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg ",
		"cn": "TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb ",
		"gn": "PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi ",
		"dn": "____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"an": "____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ",
		"en": "SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ ",
		"bn": "AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ "}

	k5 = map[string]string{ // 1426353696163
		"fn": "npFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"an": "HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"en": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ ",
		"bn": "SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ "}

	k15 = map[string]string{ // 1426353696165
		"fn": "npFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg ",
		"cn": "____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"en": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"bn": "SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ "}

	j56 = map[string]string{ // 1426353696167
		"fn": "npFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb ",
		"gn": "____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ",
		"dn": "PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ ",
		"an": "____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ ",
		"bn": "SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ "}

	k135x4 = map[string]string{ // 1426353696169
		"fn": "npFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ ",
		"cn": "____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ",
		"gn": "TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ ",
		"dn": "PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp ",
		"en": "____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"bn": "SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ "}

	j56y7 = map[string]string{ // 1426353696171
		"fn": "npFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ ",
		"gn": "____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ",
		"dn": "PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ ",
		"an": "HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ ",
		"bn": "____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ "}

	n345y7 = map[string]string{ // 1426353696172
		"fn": "npFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg ",
		"cn": "____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ",
		"gn": "TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ ",
		"dn": "____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp ",
		"en": "AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ ",
		"bn": "SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ "}

	k1j6y7 = map[string]string{ // 1426353696174
		"fn": "npFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg ",
		"cn": "____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ",
		"gn": "TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr ",
		"an": "HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp ",
		"en": "AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ ",
		"bn": "____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ "}

	k345x2 = map[string]string{ // 1426353696176
		"fn": "npFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg ",
		"cn": "UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ",
		"dn": "____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"an": "HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ ",
		"bn": "SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ "}

	k2j56y7 = map[string]string{ // 1426353696178
		"fn": "PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ",
		"an": "PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"en": "HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu ",
		"bn": "____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ "}

	j346y5 = map[string]string{ // 1426353696180
		"fn": "PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg ",
		"cn": "npCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb ",
		"gn": "____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu ",
		"bn": "AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ "}

	j3k5 = map[string]string{ // 1426353696182
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ ",
		"cn": "npCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ",
		"dn": "TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp ",
		"en": "____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ",
		"bn": "AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ "}

	n167 = map[string]string{ // 1426353696184
		"fn": "PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ ",
		"cn": "npCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ",
		"dn": "____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp ",
		"en": "____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ "}

	k2 = map[string]string{ // 1426353696186
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePb HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	k125 = map[string]string{ // 1426353696188
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	k25 = map[string]string{ // 1426353696190
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePb HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	k12 = map[string]string{ // 1426353696192
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	k12j5 = map[string]string{ // 1426353696194
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	k2x1 = map[string]string{ // 1426353696196
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	k2j5 = map[string]string{ // 1426353696198
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	k25x1 = map[string]string{ // 1426353696200
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	k2j56 = map[string]string{ // 1426353696202
		"fn": "PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb ",
		"gn": "____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"en": "HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ "}

	k257x1 = map[string]string{ // 1426353696204
		"fn": "PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg ",
		"cn": "npCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ ",
		"an": "PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu ",
		"bn": "____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	k2j5y6 = map[string]string{ // 1426353696206
		"fn": "PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb ",
		"gn": "UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ",
		"en": "HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ "}

	k25x17 = map[string]string{ // 1426353696208
		"fn": "PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg ",
		"cn": "____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr ",
		"an": "PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu ",
		"bn": "____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	k2j5x1 = map[string]string{ // 1426353696209
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ",
		"dn": "TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	k1j5 = map[string]string{ // 1426353696211
		"fn": "FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg ",
		"cn": "____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ",
		"gn": "____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ",
		"dn": "npSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr ",
		"an": "UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ ",
		"en": "TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ ",
		"bn": "PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ "}

	n67x2 = map[string]string{ // 1426353696213
		"fn": "FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg ",
		"cn": "____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ",
		"gn": "____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ",
		"dn": "____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ",
		"an": "UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp ",
		"en": "TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ ",
		"bn": "PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ "}

	k1j56y7 = map[string]string{ // 1426353696215
		"fn": "FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg ",
		"cn": "____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ",
		"gn": "____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ",
		"dn": "npSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr ",
		"an": "UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp ",
		"en": "TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ ",
		"bn": "____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ "}

	j2k5 = map[string]string{ // 1426353696217
		"fn": "CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb ",
		"gn": "____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ",
		"dn": "____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ",
		"an": "npAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp ",
		"en": "MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ "}

	n6x2 = map[string]string{ // 1426353696219
		"fn": "CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ ",
		"gn": "____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ",
		"an": "npAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ",
		"en": "MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ ",
		"bn": "____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ "}

	j25 = map[string]string{ // 1426353696221
		"fn": "CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb ",
		"gn": "____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ",
		"an": "npAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ ",
		"en": "MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ "}

	j256 = map[string]string{ // 1426353696223
		"fn": "CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb ",
		"gn": "____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ",
		"an": "____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ",
		"en": "MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ "}
)

func main() {
	p := fmt.Println

	p("\nn0")
	p(n0["fn"])
	p(n0["cn"])
	p(n0["gn"])
	p(n0["dn"])
	p(n0["an"])
	p(n0["en"])
	p(n0["bn"])

	p("\nk6")
	p(k6["fn"])
	p(k6["cn"])
	p(k6["gn"])
	p(k6["dn"])
	p(k6["an"])
	p(k6["en"])
	p(k6["bn"])

	p("\nj17")
	p(j17["fn"])
	p(j17["cn"])
	p(j17["gn"])
	p(j17["dn"])
	p(j17["an"])
	p(j17["en"])
	p(j17["bn"])

	p("\nk6x5")
	p(k6x5["fn"])
	p(k6x5["cn"])
	p(k6x5["gn"])
	p(k6x5["dn"])
	p(k6x5["an"])
	p(k6x5["en"])
	p(k6x5["bn"])

	p("\nj17y2")
	p(j17y2["fn"])
	p(j17y2["cn"])
	p(j17y2["gn"])
	p(j17y2["dn"])
	p(j17y2["an"])
	p(j17y2["en"])
	p(j17y2["bn"])

	p("\nj3")
	p(j3["fn"])
	p(j3["cn"])
	p(j3["gn"])
	p(j3["dn"])
	p(j3["an"])
	p(j3["en"])
	p(j3["bn"])

	p("\nj34k6")
	p(j34k6["fn"])
	p(j34k6["cn"])
	p(j34k6["gn"])
	p(j34k6["dn"])
	p(j34k6["an"])
	p(j34k6["en"])
	p(j34k6["bn"])

	p("\nj17k2")
	p(j17k2["fn"])
	p(j17k2["cn"])
	p(j17k2["gn"])
	p(j17k2["dn"])
	p(j17k2["an"])
	p(j17k2["en"])
	p(j17k2["bn"])

	p("\nn26y5")
	p(n26y5["fn"])
	p(n26y5["cn"])
	p(n26y5["gn"])
	p(n26y5["dn"])
	p(n26y5["an"])
	p(n26y5["en"])
	p(n26y5["bn"])

	p("\nk26x5")
	p(k26x5["fn"])
	p(k26x5["cn"])
	p(k26x5["gn"])
	p(k26x5["dn"])
	p(k26x5["an"])
	p(k26x5["en"])
	p(k26x5["bn"])

	p("\nj6")
	p(j6["fn"])
	p(j6["cn"])
	p(j6["gn"])
	p(j6["dn"])
	p(j6["an"])
	p(j6["en"])
	p(j6["bn"])

	p("\nj36")
	p(j36["fn"])
	p(j36["cn"])
	p(j36["gn"])
	p(j36["dn"])
	p(j36["an"])
	p(j36["en"])
	p(j36["bn"])

	p("\nk56")
	p(k56["fn"])
	p(k56["cn"])
	p(k56["gn"])
	p(k56["dn"])
	p(k56["an"])
	p(k56["en"])
	p(k56["bn"])

	p("\nj136y7")
	p(j136y7["fn"])
	p(j136y7["cn"])
	p(j136y7["gn"])
	p(j136y7["dn"])
	p(j136y7["an"])
	p(j136y7["en"])
	p(j136y7["bn"])

	p("\nk56x4")
	p(k56x4["fn"])
	p(k56x4["cn"])
	p(k56x4["gn"])
	p(k56x4["dn"])
	p(k56x4["an"])
	p(k56x4["en"])
	p(k56x4["bn"])

	p("\nn167x4")
	p(n167x4["fn"])
	p(n167x4["cn"])
	p(n167x4["gn"])
	p(n167x4["dn"])
	p(n167x4["an"])
	p(n167x4["en"])
	p(n167x4["bn"])

	p("\nj3k5x4")
	p(j3k5x4["fn"])
	p(j3k5x4["cn"])
	p(j3k5x4["gn"])
	p(j3k5x4["dn"])
	p(j3k5x4["an"])
	p(j3k5x4["en"])
	p(j3k5x4["bn"])

	p("\nj167y2")
	p(j167y2["fn"])
	p(j167y2["cn"])
	p(j167y2["gn"])
	p(j167y2["dn"])
	p(j167y2["an"])
	p(j167y2["en"])
	p(j167y2["bn"])

	p("\nj2")
	p(j2["fn"])
	p(j2["cn"])
	p(j2["gn"])
	p(j2["dn"])
	p(j2["an"])
	p(j2["en"])
	p(j2["bn"])

	p("\nj236")
	p(j236["fn"])
	p(j236["cn"])
	p(j236["gn"])
	p(j236["dn"])
	p(j236["an"])
	p(j236["en"])
	p(j236["bn"])

	p("\nj26")
	p(j26["fn"])
	p(j26["cn"])
	p(j26["gn"])
	p(j26["dn"])
	p(j26["an"])
	p(j26["en"])
	p(j26["bn"])

	p("\nj23")
	p(j23["fn"])
	p(j23["cn"])
	p(j23["gn"])
	p(j23["dn"])
	p(j23["an"])
	p(j23["en"])
	p(j23["bn"])

	p("\nj23k6")
	p(j23k6["fn"])
	p(j23k6["cn"])
	p(j23k6["gn"])
	p(j23k6["dn"])
	p(j23k6["an"])
	p(j23k6["en"])
	p(j23k6["bn"])

	p("\nj2y3")
	p(j2y3["fn"])
	p(j2y3["cn"])
	p(j2y3["gn"])
	p(j2y3["dn"])
	p(j2y3["an"])
	p(j2y3["en"])
	p(j2y3["bn"])

	p("\nj2k6")
	p(j2k6["fn"])
	p(j2k6["cn"])
	p(j2k6["gn"])
	p(j2k6["dn"])
	p(j2k6["an"])
	p(j2k6["en"])
	p(j2k6["bn"])

	p("\nj26y3")
	p(j26y3["fn"])
	p(j26y3["cn"])
	p(j26y3["gn"])
	p(j26y3["dn"])
	p(j26y3["an"])
	p(j26y3["en"])
	p(j26y3["bn"])

	p("\nj2k56")
	p(j2k56["fn"])
	p(j2k56["cn"])
	p(j2k56["gn"])
	p(j2k56["dn"])
	p(j2k56["an"])
	p(j2k56["en"])
	p(j2k56["bn"])

	p("\nj246y3")
	p(j246y3["fn"])
	p(j246y3["cn"])
	p(j246y3["gn"])
	p(j246y3["dn"])
	p(j246y3["an"])
	p(j246y3["en"])
	p(j246y3["bn"])

	p("\nj2k56x4")
	p(j2k56x4["fn"])
	p(j2k56x4["cn"])
	p(j2k56x4["gn"])
	p(j2k56x4["dn"])
	p(j2k56x4["an"])
	p(j2k56x4["en"])
	p(j2k56x4["bn"])

	p("\nk157x6")
	p(k157x6["fn"])
	p(k157x6["cn"])
	p(k157x6["gn"])
	p(k157x6["dn"])
	p(k157x6["an"])
	p(k157x6["en"])
	p(k157x6["bn"])

	p("\nj26y34")
	p(j26y34["fn"])
	p(j26y34["cn"])
	p(j26y34["gn"])
	p(j26y34["dn"])
	p(j26y34["an"])
	p(j26y34["en"])
	p(j26y34["bn"])

	p("\nj2k6x5")
	p(j2k6x5["fn"])
	p(j2k6x5["cn"])
	p(j2k6x5["gn"])
	p(j2k6x5["dn"])
	p(j2k6x5["an"])
	p(j2k6x5["en"])
	p(j2k6x5["bn"])

	p("\nj2k6y3")
	p(j2k6y3["fn"])
	p(j2k6y3["cn"])
	p(j2k6y3["gn"])
	p(j2k6y3["dn"])
	p(j2k6y3["an"])
	p(j2k6y3["en"])
	p(j2k6y3["bn"])

	p("\nk1j6")
	p(k1j6["fn"])
	p(k1j6["cn"])
	p(k1j6["gn"])
	p(k1j6["dn"])
	p(k1j6["an"])
	p(k1j6["en"])
	p(k1j6["bn"])

	p("\nn345")
	p(n345["fn"])
	p(n345["cn"])
	p(n345["gn"])
	p(n345["dn"])
	p(n345["an"])
	p(n345["en"])
	p(n345["bn"])

	p("\nj3k6")
	p(j3k6["fn"])
	p(j3k6["cn"])
	p(j3k6["gn"])
	p(j3k6["dn"])
	p(j3k6["an"])
	p(j3k6["en"])
	p(j3k6["bn"])

	p("\nn45y2")
	p(n45y2["fn"])
	p(n45y2["cn"])
	p(n45y2["gn"])
	p(n45y2["dn"])
	p(n45y2["an"])
	p(n45y2["en"])
	p(n45y2["bn"])

	p("\nj3k56x4")
	p(j3k56x4["fn"])
	p(j3k56x4["cn"])
	p(j3k56x4["gn"])
	p(j3k56x4["dn"])
	p(j3k56x4["an"])
	p(j3k56x4["en"])
	p(j3k56x4["bn"])

	p("\nk2j6")
	p(k2j6["fn"])
	p(k2j6["cn"])
	p(k2j6["gn"])
	p(k2j6["dn"])
	p(k2j6["an"])
	p(k2j6["en"])
	p(k2j6["bn"])

	p("\nn5y2")
	p(n5y2["fn"])
	p(n5y2["cn"])
	p(n5y2["gn"])
	p(n5y2["dn"])
	p(n5y2["an"])
	p(n5y2["en"])
	p(n5y2["bn"])

	p("\nk26")
	p(k26["fn"])
	p(k26["cn"])
	p(k26["gn"])
	p(k26["dn"])
	p(k26["an"])
	p(k26["en"])
	p(k26["bn"])

	p("\nk256")
	p(k256["fn"])
	p(k256["cn"])
	p(k256["gn"])
	p(k256["dn"])
	p(k256["an"])
	p(k256["en"])
	p(k256["bn"])

	p("")
}
