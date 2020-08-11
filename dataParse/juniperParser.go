package dataParse

import (
	"fmt"
	"regexp"
	"strings"
)

type JuniperInt struct {
	slot	string
	subslot	string
	intf	string
}

func JuniperParser(queryOut string) []SnInfo {
	var splitStr string
	var delUnuse []string

	splitStr = strings.Split(queryOut, "{master}")[0]
	eachLineSlice := strings.Split(splitStr, "\n")
	for _, line := range eachLineSlice {
		if strings.Contains(line, "SFP+") ||
			strings.Contains(line, "FPC ") ||
			strings.Contains(line, "PIC ") ||
			strings.Contains(line, "QSFP-") ||
			strings.Contains(line, "100GBASE-LR") ||
			strings.Contains(line, "100G-SR") {
			delUnuse = append(delUnuse, line)
		}
	}

	//var juniperIntfSlice []JuniperInt
	var (
		juniperFpc		string
		juniperPic		string
		juniperXcvr		string
		juniperXcvrType	string
		juniperXcvrSn	string
		juniperPicType	string
	)
	var juniperFpcs []string
	juniperFpcPic := make(map[string]string)

	for _, line := range delUnuse {
		if !strings.HasPrefix(line, " ") && strings.Contains(line, "FPC") {
			FPCReg := regexp.MustCompile(`^(FPC\s[0-9]+)\s+`)
			juniperFpc = FPCReg.FindStringSubmatch(line)[1]
			juniperFpcs = append(juniperFpcs, juniperFpc)
			continue
		}
		if strings.HasPrefix(line, "    PIC") {
			PICReg := regexp.MustCompile(`^\s+(PIC\s[0-9]+)\s+.*(OTN|SFPP|CGE)`)
			juniperPic = PICReg.FindStringSubmatch(line)[1]
			juniperPicType = PICReg.FindStringSubmatch(line)[2]
			continue
			fmt.Println(juniperPic, juniperPicType)
		}
		if strings.HasPrefix(line, "      Xcvr") {
			XcvrReg := regexp.MustCompile(`\s+(Xcvr\s[0-9]+)\s+(REV\s[0-9]+\s+[0-9]+-[0-9]+\s+|NON-JNPR\s+)([A-Z0-9]+)\s+(.*)`)
			juniperXcvr = XcvrReg.FindStringSubmatch(line)[1]
			juniperXcvrSn = XcvrReg.FindStringSubmatch(line)[3]
			juniperXcvrType = XcvrReg.FindStringSubmatch(line)[4]
			continue
			fmt.Printf("%s -- %s -- %s\n", juniperXcvr, juniperXcvrType, juniperXcvrSn)
		}
	}
	for _, fpc := range juniperFpcs {
		juniperFpcPic[fpc] = ""
	}
	fmt.Println(juniperFpcPic)
	return []SnInfo{}
}
