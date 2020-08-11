package dataParse

import (
	"regexp"
	"strings"
)

func HuaweiParser(queryOut string) []SnInfo {
	var splitStr	[]string
	var delUnuseTmp	[]string
	var delUnuse	[]string

	splitStr = splitByEmptyLine(queryOut)[1:]
	for _, val := range splitStr {
		if !strings.Contains(val, "Active_Optical") {
			delUnuseTmp = append(delUnuseTmp, val)
		}
	}
	for _, val := range delUnuseTmp {
		if strings.Contains(val, "information") {
			delUnuse = append(delUnuse, val)
		}
	}

	var snInfoSlice []SnInfo
	for idx, val := range delUnuse {
		intNameReg := regexp.MustCompile(`\s(.*?)\stransceiver\sinformation`)
		modTypeReg := regexp.MustCompile(`Transceiver\sType\s+:(.*?)\n`)
		modSnReg := regexp.MustCompile(`Serial\sNumber\s+:(.*?)\n`)

		intName := intNameReg.FindStringSubmatch(val)[1]
		modType := modTypeReg.FindStringSubmatch(val)[1]
		modSn := modSnReg.FindStringSubmatch(val)[1]

		snInfoSlice = append(snInfoSlice, SnInfo{
			Id: idx,
			IntName: intName,
			ModType: modType,
			ModSn: modSn,
		})
	}
	return snInfoSlice
}
