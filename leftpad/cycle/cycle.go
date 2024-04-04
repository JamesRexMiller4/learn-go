package cycle

import (
	"_/Users/jamesmiller/Desktop/learn-go/leftpad"
)

var DEFAULT_CHAR = ' '

func FormatDouble(s string, i int) string {
	return leftpad.Format(s+s, i)
}
