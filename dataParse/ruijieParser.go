package dataParse

import (
	"regexp"
	"strings"
)

func RuijieParser(queryOut string) []SnInfo {
	var splitStr	[]string
	var delUnuse	[]string
	splitStr = strings.Split(queryOut, "===Interface ")[1:]
	for _, val := range splitStr {
		if !(strings.Contains(val, "absent") || strings.Contains(val, "G-Active")) {
			delUnuse = append(delUnuse, val)
		}
	}
	var snInfoSlice []SnInfo
	for idx, val := range delUnuse {
		intNameReg	:= regexp.MustCompile(`(.*?)=+\n`)
		modTypeReg	:= regexp.MustCompile(`Transceiver\sType\s+:\s+(.*?)\n`)
		modSnReg	:= regexp.MustCompile(`Vendor\sSerial\sNumber\s+:\s(.*?)\n`)

		intName	:= intNameReg.FindStringSubmatch(val)[1]
		modType	:= modTypeReg.FindStringSubmatch(val)[1]
		modSn	:= modSnReg.FindStringSubmatch(val)[1]

		snInfoSlice = append(snInfoSlice, SnInfo{
			Id: idx,
			IntName: intName,
			ModType: modType,
			ModSn: modSn,
		})
	}
	return snInfoSlice
}
