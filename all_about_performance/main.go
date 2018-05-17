package main

import (
	"strconv"
	"strings"
)

// format takes 2 arguments and output a formatted string
// sample: format(1,2.0) => "formated-int(1)-float(2.00)" the float number shall always have 2 digits after decimal point.
func format(intValue int64, floatValue float64) string {
	var buffer strings.Builder
	buffer.WriteString("formatted-int(" + strconv.FormatInt(intValue, 10) + ")-float(" + strconv.FormatFloat(floatValue, 'f', 2, 64) + ")")
	return buffer.String()
}
