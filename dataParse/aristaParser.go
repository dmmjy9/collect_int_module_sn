package dataParse

import (
	"regexp"
	"strings"
)

func AristaParser(queryOut string) []SnInfo {
	var splitStr []string
	var delUnuse []string
	splitStr = strings.Split(queryOut, "Port ")
	for _, val := range splitStr {
		if !strings.Contains(val, "cable") {
			delUnuse = append(delUnuse, val)
		}
	}
	delUnuse = delUnuse[1:]

	var snInfoSlice []SnInfo
	for idx, val := range delUnuse {
		intNameReg	:= regexp.MustCompile(`(Ethernet.*?)\n`)
		modTypeReg	:= regexp.MustCompile(`Transceiver\s+(.*?)\s+`)
		modSnReg	:= regexp.MustCompile(`Transceiver\sSN\s+(.*?)\s+`)

		intName	:= intNameReg.FindStringSubmatch(val)[1]
		modType	:= modTypeReg.FindStringSubmatch(val)[1]
		modSn	:= modSnReg.FindStringSubmatch(val)[1]
		snInfoSlice = append(snInfoSlice, SnInfo{
			Id:      idx,
			IntName: intName,
			ModType: modType,
			ModSn:   modSn,
		})
	}

	return snInfoSlice
}
