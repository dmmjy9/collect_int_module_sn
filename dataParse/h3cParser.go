package dataParse

import (
	"regexp"
	"strings"
)

func transInfoParser(transInfo string) map[string]string {
	transInfoMap := make(map[string]string)
	transInfoLines := strings.Split(transInfo, "\n")
	var intName string

	for _, val := range transInfoLines {
		if strings.Contains(val, "M-Giga") {
			continue
		}
		if !strings.HasPrefix(val, " ") && strings.Contains(val, " transceiver information:") {
			intName = strings.Split(val, " transceiver information:")[0]
			continue
		} else {
			if strings.Contains(val, "absent") || strings.Contains(val, "AOC") {
				delete(transInfoMap, intName)
				continue
			}
			if strings.Contains(val, "Transceiver Type") {
				n := strings.Split(val, ": ")[1]
				transInfoMap[intName] = n
				continue
			}
		}
	}
	return transInfoMap
}

func transManuParser(transManu string, interfaces []string) map[string]string {
	transManuMap := make(map[string]string)

	for _, intf := range interfaces {
		modSnReg := regexp.MustCompile(intf+`.*\n.*?:\s(.*?)\n`)
		modSn := modSnReg.FindStringSubmatch(transManu)
		if len(modSn) != 0 {
			transManuMap[intf] = modSn[1]
		}
	}
	return transManuMap
}

func H3CParser(queryOut string) []SnInfo {
	var snInfoSlice []SnInfo
	var intfSlice []string
	allInfo := strings.Split(
		strings.Split(queryOut, "display transceiver interface")[1], "quit")[0]
	transInfo := strings.Split(allInfo, "display transceiver manuinfo interface")[0]
	transManu := strings.Split(allInfo, "display transceiver manuinfo interface")[1]
	transInfoMap := transInfoParser(transInfo)
	for intf, _ := range transInfoMap {
		intfSlice = append(intfSlice, intf)
	}
	transManuMap := transManuParser(transManu, intfSlice)

	id := 0
	for intf, info := range transInfoMap {
		snInfoSlice = append(snInfoSlice, SnInfo{
			Id: id,
			IntName: intf,
			ModType: info,
			ModSn: transManuMap[intf],
		})
		id++
	}

	return snInfoSlice
}
