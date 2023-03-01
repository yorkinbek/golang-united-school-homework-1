package hard

import (
	"io"
	"strings"
)

func GoalParsers(strReader *strings.Reader) string {
	var result string
	for {
		r, _, err := strReader.ReadRune()
		if err == io.EOF {
			break
		}
		switch r {
		case 'G':
			result += "G"
		case '(':
			r, _, err := strReader.ReadRune()
			if err == io.EOF {
				break
			}
			if r == ')' {
				result += "o"
			} else {
				strReader.UnreadRune()
				str, _, _ := strReader.ReadRune()
				str, _, _ = strReader.ReadRune()
				if str == 'l' {
					result += "al"
				}
			}
		}
	}
	return result
}
