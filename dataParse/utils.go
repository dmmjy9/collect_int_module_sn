package dataParse

import "regexp"

type SnInfo struct {
	Id			int
	IntName		string
	ModType		string
	ModSn		string
}

func splitByEmptyLine(str string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(str, "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
}
