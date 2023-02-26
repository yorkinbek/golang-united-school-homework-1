package homework

import (
	"io/ioutil"
	"strings"
)

func SeekTillHalfOfString(strReader *strings.Reader) string {

	strLen := strReader.Len()

	midPoint := strLen / 2

	_, err := strReader.Seek(int64(midPoint), 0)
	if err != nil {
		panic(err)
	}

	remainingLen := strLen - midPoint
	buf := make([]byte, remainingLen)
	_, err = strReader.Read(buf)
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func ReaderSplit(strReader *strings.Reader, n int) []string {
	var chunks []string
	buf, err := ioutil.ReadAll(strReader)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i += n {
		end := i + n
		if end > len(buf) {
			end = len(buf)
		}
		chunks = append(chunks, string(buf[i:end]))
	}

	return chunks
}
