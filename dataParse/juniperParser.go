package dataParse

import (
	"regexp"
	"strings"
)

func JuniperParser(queryOut string) []SnInfo {
	eachLineSlice := strings.Split(queryOut, "\n")
	var snInfoSlice []SnInfo
	idx := 0

	var FPCNow, PICNow, XcvrNow, TypeNow string

	for _, line := range eachLineSlice {
		if strings.Contains(line, "FPC") {
			FPCReg := regexp.MustCompile(`^FPC\s([0-9]+)`)
			FPCNow = FPCReg.FindStringSubmatch(line)[1]
		} else if strings.Contains(line, "PIC") {
			PICReg := regexp.MustCompile(`^\s+PIC\s([0-9]+)`)
			TypeReg := regexp.MustCompile(`(OTN$|SFPP$|-CGE$)`)
			PICNow = PICReg.FindStringSubmatch(line)[1]
			TypeNow = func(reg *regexp.Regexp) string {
				matchType := reg.FindStringSubmatch(line)[1]
				if matchType == "OTN" || matchType == "-CGE" {
					return "et"
				} else if matchType == "SFPP" {
					return "xe"
				}
				return ""
			}(TypeReg)
		} else if strings.Contains(line, "Xcvr") {
			XcvrReg := regexp.MustCompile(`^\s+Xcvr\s([0-9]+)`)
			XcvrNow = XcvrReg.FindStringSubmatch(line)[1]
			idx++

			intName := TypeNow + "-" + FPCNow + "/" + PICNow + "/" + XcvrNow
			modTypeReg := regexp.MustCompile(`([A-Za-z0-9\+]+\-[A-Za-z0-9\+]+\-[A-Za-z0-9\+]+|[A-Za-z0-9\+]+\-[A-Za-z0-9\+]+\-[A-Za-z0-9\+]+\-[A-Za-z0-9\+]+)`)
			modType := modTypeReg.FindStringSubmatch(line)[1]
			modSnReg := regexp.MustCompile(`([A-Za-z0-9]+)\s+(SFP|QSFP|CFP|100GB)`)
			modSn := modSnReg.FindStringSubmatch(line)[1]

			snInfoSlice = append(snInfoSlice, SnInfo{
				Id:      idx,
				IntName: intName,
				ModType: modType,
				ModSn:   modSn,
			})
		} else {
			continue
		}
	}

	return snInfoSlice
}
