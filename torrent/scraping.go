package torrent

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func clearSpaces(raw string) string {
	txt := strings.Replace(raw, "\n", "", -1)
	txt = strings.Replace(txt, "\t", "  ", -1)
	txt = strings.Replace(txt, "  ", " ", -1)
	return txt
}

func fixMonths(str string) string {
	months := map[string]string{
		"Янв": "Jan",
		"Фев": "Feb",
		"Mar": "Mar",
		"Апр": "Apr",
		"Май": "May",
		"Июн": "Jun",
		"Июл": "Jul",
		"Авг": "Aug",
		"Сен": "Sep",
		"Окт": "Oct",
		"Ноя": "Nov",
		"Дек": "Dec",
	}
	for r, en := range months {
		str = strings.Replace(str, r, en, -1)
	}
	return str
}

func formatTime(str string) time.Time {
	//7-Апр-20 00:06
	str = strings.Trim(str, " \t\n\r")
	str = strings.Replace(str, "  ", " ", -1)
	str = strings.Replace(str, "  ", " ", -1)
	str = fixMonths(str)
	t, err := time.Parse("2-Jan-06 15:04", str)
	//t, err := time.Parse("2-Jan-20 17:35:00", str)
	if err != nil {
		log.Errorf("Error while parsing time string: %s\t %v\n", str, err)
		return time.Now()
	}
	return t
}

func extractAttr(uri string, param string) string {
	furl, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	val := furl.Query().Get(param)
	return val
}

func stripToNumber(str string) string {
	chars := "0123456789.,"
	var validChars []rune
	for _, c := range []rune(str) {
		if strings.ContainsRune(chars, c) {
			validChars = append(validChars, c)
		}
	}
	str = string(validChars)
	return str
}

func sizeStrToBytes(str string) uint64 {
	str = strings.ToLower(str)
	str = clearSpaces(str)
	multiplier := 1
	if strings.Contains(str, "gb") {
		multiplier = 1028 * 1028 * 1028
	} else if strings.Contains(str, "mb") {
		multiplier = 1028 * 1028
	} else if strings.Contains(str, "kb") {
		multiplier = 1028
	}
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "gb", "", -1)
	str = strings.Replace(str, "mb", "", -1)
	str = strings.Replace(str, "kb", "", -1)
	chars := "1203456789.,"
	var validChars []rune
	for _, c := range []rune(str) {
		if strings.ContainsRune(chars, c) {
			validChars = append(validChars, c)
		}
	}
	str = string(validChars)
	flt, _ := strconv.ParseFloat(str, 32)
	flt = flt * float64(multiplier)
	return uint64(flt)
}