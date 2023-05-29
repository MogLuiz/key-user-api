package entity

import "strconv"

func ParseInt(s string) int {
	if s == "" {
		return 0
	}
	convertedString, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return convertedString
}
