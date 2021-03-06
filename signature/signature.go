package signature

var (
	N0 = map[string]string{ // 1426353696131
		"fn": "PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ CuHg ",
		"cn": "HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ FePb ",
		"gn": "AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ CuHg PbFe ____ ",
		"dn": "SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ ",
		"an": "AgAu ____ FePb HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ ",
		"en": "CuHg PbFe ____ AuAg ____ AgAu ____ FePb HgCu ____ SnSn ____ ",
		"bn": "FePb HgCu ____ SnSn ____ CuHg PbFe ____ AuAg ____ AgAu ____ "}

	K6 = map[string]string{ // 1426353696133
		"fn": "PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"cn": "HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi ",
		"gn": "AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ ",
		"an": "____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ ",
		"bn": "FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe "}

	J17 = map[string]string{ // 1426353696143
		"fn": "PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb ",
		"cn": "____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"gn": "AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ ",
		"an": "AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ ",
		"bn": "FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe "}

	K6X5 = map[string]string{ // 1426353696146
		"fn": "PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"cn": "HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi ",
		"gn": "____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"dn": "SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ ",
		"an": "AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ ",
		"en": "CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ ",
		"bn": "FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe "}

	J17Y2 = map[string]string{ // 1426353696150
		"fn": "PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb ",
		"cn": "HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"gn": "AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ ",
		"dn": "____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ",
		"an": "AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ ",
		"en": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ ",
		"bn": "FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe "}

	J3 = map[string]string{ // 1426353696153
		"fn": "PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ ",
		"cn": "HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr ",
		"gn": "AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ ",
		"dn": "SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ ",
		"an": "AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ ",
		"en": "____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ",
		"bn": "FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ "}

	J34K6 = map[string]string{ // 1426353696155
		"fn": "____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ",
		"cn": "HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn ",
		"gn": "AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ ",
		"dn": "SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ ",
		"an": "____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ",
		"en": "CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe ",
		"bn": "FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu "}

	K2J17 = map[string]string{ // 1578340594509
		"fn": "PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi ",
		"cn": "____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ",
		"gn": "AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ ",
		"dn": "____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ",
		"an": "AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ ",
		"en": "CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe ",
		"bn": "FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu "}

	N26Y5 = map[string]string{ // 1426353696159
		"fn": "PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi ",
		"cn": "____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"gn": "____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ",
		"dn": "SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ ",
		"an": "AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ ",
		"en": "CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe ",
		"bn": "FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu "}

	K26X5 = map[string]string{ // 1426353696161
		"fn": "PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi ",
		"cn": "HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"gn": "____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ",
		"dn": "____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ",
		"an": "AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ ",
		"en": "CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe ",
		"bn": "FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu "}

	J6 = map[string]string{ // 1426353696163
		"fn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"gn": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"dn": "SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ ",
		"an": "____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ",
		"en": "CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ "}

	J36 = map[string]string{ // 1426353696165
		"fn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"cn": "HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"gn": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ ",
		"dn": "SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ ",
		"an": "____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ",
		"bn": "FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ "}

	K56 = map[string]string{ // 1426353696167
		"fn": "PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp ",
		"gn": "____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"dn": "SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ",
		"en": "CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg "}

	J136Y7 = map[string]string{ // 1426353696169
		"fn": "PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ ",
		"cn": "____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ ",
		"dn": "SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ ",
		"an": "AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ",
		"bn": "FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ "}

	K56X4 = map[string]string{ // 1426353696171
		"fn": "____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"cn": "HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp ",
		"gn": "AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ ",
		"dn": "SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ",
		"en": "CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ ",
		"bn": "FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg "}

	N167X4 = map[string]string{ // 1426353696172
		"fn": "PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ ",
		"cn": "HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ ",
		"dn": "____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ",
		"an": "AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe ",
		"en": "____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ",
		"bn": "FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ "}

	J3K5X4 = map[string]string{ // 1426353696174
		"fn": "____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"cn": "HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp ",
		"gn": "AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ ",
		"dn": "SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ",
		"an": "AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe ",
		"en": "____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ",
		"bn": "FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ "}

	J167Y2 = map[string]string{ // 1426353696176
		"fn": "PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr ",
		"cn": "HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp ",
		"gn": "AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ ",
		"dn": "____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ",
		"an": "____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ",
		"en": "CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ ",
		"bn": "FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg "}

	J2K56X4 = map[string]string{ // 1426353696178
		"fn": "____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu ",
		"gn": "AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ ",
		"dn": "____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ",
		"en": "CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn "}

	K157X6 = map[string]string{ // 1426353696180
		"fn": "PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"cn": "HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu ",
		"gn": "____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ",
		"dn": "SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ",
		"en": "CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ "}

	K1J6 = map[string]string{ // 1426353696182
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp ",
		"cn": "____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ ",
		"dn": "SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe ",
		"an": "____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ ",
		"bn": "FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ "}

	N345 = map[string]string{ // 1426353696184
		"fn": "PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp ",
		"cn": "____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ",
		"gn": "AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ",
		"an": "____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ ",
		"bn": "FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn "}

	J2 = map[string]string{ // 1426353696186
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ",
		"en": "CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ "}

	J26 = map[string]string{ // 1426353696188
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ",
		"en": "CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ "}

	J236 = map[string]string{ // 1426353696190
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePb HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ",
		"en": "____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ "}

	J23 = map[string]string{ // 1426353696192
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ "}

	J23K6 = map[string]string{ // 1426353696194
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ ",
		"cn": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ ",
		"dn": "____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ",
		"bn": "FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn "}

	J2Y3 = map[string]string{ // 1426353696196
		"fn": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"gn": "AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ",
		"an": "AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ "}

	J2K6 = map[string]string{ // 1426353696198
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ",
		"en": "CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn "}

	J26Y3 = map[string]string{ // 1426353696200
		"fn": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ",
		"en": "____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ "}

	J2K56 = map[string]string{ // 1426353696202
		"fn": "PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu ",
		"gn": "____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ",
		"en": "CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn "}

	J246Y3 = map[string]string{ // 1426353696204
		"fn": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"cn": "HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ ",
		"dn": "SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ",
		"en": "CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ "}

	J26Y34 = map[string]string{ // 1426353696206
		"fn": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ",
		"cn": "HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu ",
		"gn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ ",
		"dn": "SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe ",
		"an": "____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ",
		"en": "____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ",
		"bn": "FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ "}

	J2K6X5 = map[string]string{ // 1426353696208
		"fn": "PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"cn": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu ",
		"gn": "____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ",
		"dn": "____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ",
		"an": "AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ ",
		"en": "CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ ",
		"bn": "FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn "}

	J2K6Y3 = map[string]string{ // 1426353696209
		"fn": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ ",
		"cn": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"gn": "AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ ",
		"dn": "SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ",
		"an": "____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ",
		"en": "____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ",
		"bn": "FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn "}

	J3K6 = map[string]string{ // 1426353696211
		"fn": "PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ",
		"cn": "HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe ",
		"gn": "AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ ",
		"dn": "SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ",
		"en": "____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ",
		"bn": "FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg "}

	N45Y2 = map[string]string{ // 1426353696213
		"fn": "PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ ",
		"cn": "HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe ",
		"gn": "AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ ",
		"dn": "____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ",
		"en": "____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ",
		"bn": "FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg "}

	J3K56X4 = map[string]string{ // 1426353696215
		"fn": "____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ",
		"cn": "HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe ",
		"gn": "AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ ",
		"dn": "SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ",
		"an": "____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ",
		"en": "____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ",
		"bn": "FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg "}

	K2J6 = map[string]string{ // 1426353696217
		"fn": "PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu ",
		"gn": "AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ",
		"an": "____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ",
		"en": "CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ "}

	N5Y2 = map[string]string{ // 1426353696219
		"fn": "____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu ",
		"gn": "AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ ",
		"dn": "____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ",
		"en": "CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb "}

	K26 = map[string]string{ // 1426353696221
		"fn": "PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu ",
		"gn": "AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ",
		"en": "CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb "}

	K256 = map[string]string{ // 1426353696223
		"fn": "PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ",
		"cn": "HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu ",
		"gn": "____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ",
		"dn": "____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ",
		"an": "____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ",
		"en": "CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"bn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb "}

	J5 = map[string]string{ // 1426353696133
		"fn": "TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg ",
		"cn": "PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb ",
		"gn": "____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"dn": "AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ ",
		"en": "AgHg TiFe FeTi ____ ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ____ ____ SnAu ____ "}

	K34 = map[string]string{ // 1426353696143
		"fn": "TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ ",
		"cn": "PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"gn": "HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi ",
		"dn": "AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ ",
		"en": "____ TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ ____ TiFe FeTi HgAg ____ SnAu ____ "}

	J5Y6 = map[string]string{ // 1426353696146
		"fn": "TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg ",
		"cn": "PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb ",
		"gn": "HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi ",
		"dn": "AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ ",
		"an": "____ ____ CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ",
		"en": "AgHg TiFe FeTi HgAg ____ ____ ____ CuPb PbCu ____ AuSn ____ ",
		"bn": "CuPb PbCu ____ AuSn ____ AgHg TiFe FeTi HgAg ____ ____ ____ "}

	K34X2 = map[string]string{ // 1426353696150
		"fn": "TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg ",
		"cn": "PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb ",
		"gn": "HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi ",
		"dn": "____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ",
		"an": "SnAu ____ CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ ",
		"en": "AgHg TiFe FeTi HgAg ____ SnAu ____ CuPb PbCu ____ ____ ____ ",
		"bn": "CuPb PbCu ____ ____ ____ AgHg TiFe FeTi HgAg ____ SnAu ____ "}

	K1 = map[string]string{ // 1426353696153
		"fn": "UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg ",
		"cn": "____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ",
		"gn": "PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ ",
		"dn": "HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr ",
		"an": "AuAu ____ AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ ",
		"en": "SnHg UrFe ____ PbAg ____ AuAu ____ AgPb ____ FeUr HgSn ____ ",
		"bn": "AgPb ____ FeUr HgSn ____ SnHg UrFe ____ PbAg ____ AuAu ____ "}

	K17J5 = map[string]string{ // 1426353696155
		"fn": "MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg ",
		"cn": "TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ ",
		"gn": "____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ",
		"dn": "HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn ",
		"an": "AuAu ____ ____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ ",
		"en": "SnHg MnFe CuTi ____ ____ AuAu ____ ____ TiCu FeMn HgSn ____ ",
		"bn": "____ TiCu FeMn HgSn ____ SnHg MnFe CuTi ____ ____ AuAu ____ "}

	J2K34 = map[string]string{ // 1426353696157
		"fn": "MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ ",
		"cn": "TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb ",
		"gn": "PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi ",
		"dn": "____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ",
		"an": "AuAu ____ AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ ",
		"en": "____ MnFe CuTi PbAg ____ AuAu ____ AgPb TiCu FeMn ____ ____ ",
		"bn": "AgPb TiCu FeMn ____ ____ ____ MnFe CuTi PbAg ____ AuAu ____ "}

	N25X6 = map[string]string{ // 1426353696159
		"fn": "MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ ",
		"cn": "TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb ",
		"gn": "PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi ",
		"dn": "HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"an": "____ ____ AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ",
		"en": "____ MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn HgSn ____ ",
		"bn": "AgPb TiCu FeMn HgSn ____ ____ MnFe CuTi PbAg ____ ____ ____ "}

	J25Y6 = map[string]string{ // 1426353696161
		"fn": "MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg ",
		"cn": "TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb ",
		"gn": "PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi ",
		"dn": "____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ",
		"an": "____ ____ AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ",
		"en": "SnHg MnFe CuTi PbAg ____ ____ ____ AgPb TiCu FeMn ____ ____ ",
		"bn": "AgPb TiCu FeMn ____ ____ SnHg MnFe CuTi PbAg ____ ____ ____ "}

	K5 = map[string]string{ // 1426353696163
		"fn": "npFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"an": "HgAu ____ SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"en": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb UrCu ____ PbSn ____ ",
		"bn": "SnPb UrCu ____ PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ "}

	K15 = map[string]string{ // 1426353696165
		"fn": "npFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg ",
		"cn": "____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp ",
		"en": "AuHg NpFe ____ ____ FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"bn": "SnPb ____ CuUr PbSn ____ AuHg NpFe ____ ____ FeNp HgAu ____ "}

	J56 = map[string]string{ // 1426353696167
		"fn": "npFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb ",
		"gn": "____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ",
		"dn": "PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ ",
		"an": "____ ____ SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp ____ ____ SnPb UrCu ____ PbSn ____ ",
		"bn": "SnPb UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ____ ____ "}

	K135X4 = map[string]string{ // 1426353696169
		"fn": "npFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ ",
		"cn": "____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ",
		"gn": "TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ ",
		"dn": "PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp ",
		"en": "____ NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr PbSn ____ ",
		"bn": "SnPb ____ CuUr PbSn ____ ____ NpFe ____ TiAg FeNp HgAu ____ "}

	J56Y7 = map[string]string{ // 1426353696171
		"fn": "npFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg ",
		"cn": "UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ ",
		"gn": "____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ",
		"dn": "PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ ",
		"an": "HgAu ____ ____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp HgAu ____ ____ UrCu ____ PbSn ____ ",
		"bn": "____ UrCu ____ PbSn ____ AuHg NpFe AgTi ____ FeNp HgAu ____ "}

	N345Y7 = map[string]string{ // 1426353696172
		"fn": "npFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg ",
		"cn": "____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ",
		"gn": "TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ ",
		"dn": "____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ",
		"an": "HgAu ____ SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp ",
		"en": "AuHg NpFe ____ TiAg FeNp HgAu ____ SnPb ____ CuUr ____ ____ ",
		"bn": "SnPb ____ CuUr ____ ____ AuHg NpFe ____ TiAg FeNp HgAu ____ "}

	K1J6Y7 = map[string]string{ // 1426353696174
		"fn": "npFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg ",
		"cn": "____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ",
		"gn": "TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ ",
		"dn": "PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr ",
		"an": "HgAu ____ ____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp ",
		"en": "AuHg NpFe ____ TiAg FeNp HgAu ____ ____ ____ CuUr PbSn ____ ",
		"bn": "____ ____ CuUr PbSn ____ AuHg NpFe ____ TiAg FeNp HgAu ____ "}

	K345X2 = map[string]string{ // 1426353696176
		"fn": "npFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg ",
		"cn": "UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb ",
		"gn": "____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ",
		"dn": "____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ",
		"an": "HgAu ____ SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp ",
		"en": "AuHg NpFe AgTi ____ FeNp HgAu ____ SnPb UrCu ____ ____ ____ ",
		"bn": "SnPb UrCu ____ ____ ____ AuHg NpFe AgTi ____ FeNp HgAu ____ "}

	K2J56Y7 = map[string]string{ // 1426353696178
		"fn": "PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ",
		"an": "PbAu ____ ____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"en": "HgHg PuFe SnTi ____ CuNp PbAu ____ ____ NpCu ____ ____ FePu ",
		"bn": "____ NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp PbAu ____ "}

	J346Y5 = map[string]string{ // 1426353696180
		"fn": "PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg ",
		"cn": "npCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb ",
		"gn": "____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp ____ ____ AuPb NpCu ____ TiSn FePu ",
		"bn": "AuPb NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ____ ____ "}

	J3K5 = map[string]string{ // 1426353696182
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ ",
		"cn": "npCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ",
		"dn": "TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp ",
		"en": "____ PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ TiSn FePu ",
		"bn": "AuPb NpCu ____ TiSn FePu ____ PuFe ____ ____ CuNp PbAu ____ "}

	N167 = map[string]string{ // 1426353696184
		"fn": "PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ ",
		"cn": "npCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ",
		"dn": "____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp ",
		"en": "____ PuFe SnTi ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu ____ PuFe SnTi ____ CuNp PbAu ____ "}

	K2 = map[string]string{ // 1426353696186
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePb HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	K125 = map[string]string{ // 1426353696188
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	K25 = map[string]string{ // 1426353696190
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePb HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	K12 = map[string]string{ // 1426353696192
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ ",
		"dn": "____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	K12J5 = map[string]string{ // 1426353696194
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg ",
		"cn": "____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ",
		"an": "PbAu ____ AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ AgUr ____ FePu ",
		"bn": "AuPb ____ AgUr ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	K2X1 = map[string]string{ // 1426353696196
		"fn": "PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ",
		"gn": "UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ ",
		"en": "HgHg PuFe ____ UrAg ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ UrAg ____ PbAu ____ "}

	K2J5 = map[string]string{ // 1426353696198
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ",
		"an": "PbAu ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	K25X1 = map[string]string{ // 1426353696200
		"fn": "PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ",
		"gn": "____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	K2J56 = map[string]string{ // 1426353696202
		"fn": "PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb ",
		"gn": "____ CuNp ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ",
		"en": "HgHg PuFe SnTi ____ CuNp ____ ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ____ CuNp ____ ____ "}

	K257X1 = map[string]string{ // 1426353696204
		"fn": "PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg ",
		"cn": "npCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ ",
		"an": "PbAu ____ ____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ NpCu ____ TiSn FePu ",
		"bn": "____ NpCu ____ TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	K2J5Y6 = map[string]string{ // 1426353696206
		"fn": "PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg ",
		"cn": "npCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb ",
		"gn": "UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi ",
		"dn": "____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ",
		"an": "____ ____ AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ",
		"en": "HgHg PuFe SnTi UrAg ____ ____ ____ AuPb NpCu ____ ____ FePu ",
		"bn": "AuPb NpCu ____ ____ FePu HgHg PuFe SnTi UrAg ____ ____ ____ "}

	K25X17 = map[string]string{ // 1426353696208
		"fn": "PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg ",
		"cn": "____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ",
		"gn": "____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ",
		"dn": "TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr ",
		"an": "PbAu ____ ____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp ",
		"en": "HgHg PuFe ____ ____ CuNp PbAu ____ ____ ____ AgUr TiSn FePu ",
		"bn": "____ ____ AgUr TiSn FePu HgHg PuFe ____ ____ CuNp PbAu ____ "}

	K2J5X1 = map[string]string{ // 1426353696209
		"fn": "PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg ",
		"cn": "____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ",
		"gn": "____ ____ PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ",
		"dn": "TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ ",
		"an": "PbAu ____ AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ ",
		"en": "HgHg PuFe SnTi ____ ____ PbAu ____ AuPb ____ ____ TiSn FePu ",
		"bn": "AuPb ____ ____ TiSn FePu HgHg PuFe SnTi ____ ____ PbAu ____ "}

	K1J5 = map[string]string{ // 1426353696211
		"fn": "FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg ",
		"cn": "____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ",
		"gn": "____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ",
		"dn": "npSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr ",
		"an": "UrAu ____ PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ ",
		"en": "TiHg FeFe HgTi ____ ____ UrAu ____ PbPb ____ AuUr NpSn ____ ",
		"bn": "PbPb ____ AuUr NpSn ____ TiHg FeFe HgTi ____ ____ UrAu ____ "}

	N67X2 = map[string]string{ // 1426353696213
		"fn": "FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg ",
		"cn": "____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ",
		"gn": "____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ",
		"dn": "____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ",
		"an": "UrAu ____ PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp ",
		"en": "TiHg FeFe HgTi ____ SnNp UrAu ____ PbPb ____ AuUr ____ ____ ",
		"bn": "PbPb ____ AuUr ____ ____ TiHg FeFe HgTi ____ SnNp UrAu ____ "}

	K1J56Y7 = map[string]string{ // 1426353696215
		"fn": "FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg ",
		"cn": "____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ",
		"gn": "____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ",
		"dn": "npSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr ",
		"an": "UrAu ____ ____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp ",
		"en": "TiHg FeFe HgTi ____ SnNp UrAu ____ ____ ____ AuUr NpSn ____ ",
		"bn": "____ ____ AuUr NpSn ____ TiHg FeFe HgTi ____ SnNp UrAu ____ "}

	J2K5 = map[string]string{ // 1426353696217
		"fn": "CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb ",
		"gn": "____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ",
		"dn": "____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ",
		"an": "npAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp ",
		"en": "MnHg CuFe ____ ____ AuNp NpAu ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe ____ ____ AuNp NpAu ____ "}

	N6X2 = map[string]string{ // 1426353696219
		"fn": "CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ ",
		"gn": "____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ",
		"an": "npAu ____ ____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ",
		"en": "MnHg CuFe PbTi ____ AuNp NpAu ____ ____ FeCu HgMn ____ ____ ",
		"bn": "____ FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp NpAu ____ "}

	J25 = map[string]string{ // 1426353696221
		"fn": "CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb ",
		"gn": "____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ",
		"an": "npAu ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ ",
		"en": "MnHg CuFe PbTi ____ ____ NpAu ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ ____ NpAu ____ "}

	J256 = map[string]string{ // 1426353696223
		"fn": "CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg ",
		"cn": "FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb ",
		"gn": "____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ",
		"dn": "____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ",
		"an": "____ ____ TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ",
		"en": "MnHg CuFe PbTi ____ AuNp ____ ____ TiPb FeCu HgMn ____ ____ ",
		"bn": "TiPb FeCu HgMn ____ ____ MnHg CuFe PbTi ____ AuNp ____ ____ "}
)
