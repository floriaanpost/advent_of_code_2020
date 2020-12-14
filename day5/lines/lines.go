package lines

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// MustParse returns all lines in a file split by the separator
func MustParse(filename string, separator string) []string {
	buf, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	splitByToken := func(separator string) func([]byte, bool) (int, []byte, error) {
		return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if i := strings.Index(string(data), separator); i >= 0 {
				return i + len(separator), data[0:i], nil
			}
			if atEOF {
				return len(data), data, nil
			}
			return
		}
	}

	snl := bufio.NewScanner(buf)
	snl.Split(splitByToken(separator))
	var lines []string
	for snl.Scan() {
		lines = append(lines, snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}
