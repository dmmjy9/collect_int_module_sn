package dataParse

import (
	"regexp"
	"strings"
)

func CiscoParser(queryOut string) []SnInfo {
	var splitStr []string
	var delUnuse []string
	splitStr = strings.Split(queryOut, "cisco extended id number")
	for _, val := range splitStr {
		if !(strings.Contains(val, "cable") || strings.Contains(val, "not present")) {
			delUnuse = append(delUnuse, val)
		}
	}

	var snInfoSlice []SnInfo
	for idx, val := range delUnuse {
		var intName, modType, modSn string
		intNameReg := regexp.MustCompile(`(Ethernet.*?)\n`)
		modTypeReg := regexp.MustCompile(`type\sis\s(.*?)\n`)
		modSnReg := regexp.MustCompile(`serial\snumber\sis\s(.*?)\s+`)

		intName = intNameReg.FindStringSubmatch(val)[1]
		modType = modTypeReg.FindStringSubmatch(val)[1]
		if strings.Contains(modType, "(Passive)") {
			modType = strings.Split(modType, "(Passive)")[0]
		}
		modSn = modSnReg.FindStringSubmatch(val)[1]

		snInfoSlice = append(snInfoSlice, SnInfo{
			Id:      idx,
			IntName: intName,
			ModType: modType,
			ModSn:   modSn,
		})
	}

	return snInfoSlice
}
