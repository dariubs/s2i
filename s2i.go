package s2i

import (
	"strconv"
)

// Parse : string to int64
func Parse(strnum string, defaultnum interface{}) int64 {
	num, err := strconv.ParseInt(strnum, 10, 64)
	if err != nil {
		return int64(defaultnum.(int))
	}

	return num
}

// ParseUint : string to uint
func ParseUint(strnum string, defaultnum interface{}) uint {
	return uint(Parse(strnum, defaultnum))
}

// ParseInt : string to int
func ParseInt(strnum string, defaultnum interface{}) int {
	return int(Parse(strnum, defaultnum))
}

// ParseInt64 : string to int64
func ParseInt64(strnum string, defaultnum interface{}) int64 {
	num, err := strconv.ParseInt(strnum, 10, 64)
	if err != nil {
		return int64(defaultnum.(int))
	}
	return num
}
