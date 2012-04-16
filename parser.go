package configparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Section map[string]string
type ConfigFile map[string]Section

func isComment(l string) bool {
	return len(l) > 0 && l[0] == '#'
}

func isSectionHeader(l string) bool {
	return len(l) > 2 && l[0] == '[' && l[len(l)-1] == ']'
}

func isContinuation(line string) bool {
	return line[0] == ' ' || line[0] == '\t'
}

func ParseString(s string) (config ConfigFile, err error) {
	buf := bytes.NewBufferString(s)
	return Parse(buf)
}

func ParseFile(name string) (config ConfigFile, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(f)
}

func Parse(r io.Reader) (config ConfigFile, err error) {
	rd := bufio.NewReader(r)

	config = make(ConfigFile)

	var curSection Section

	var curKey string

	var line string

	lineNum := 0
	for line, err = rd.ReadString('\n'); err == nil; line, err = rd.ReadString('\n') {
		lineNum += 1
		tl := strings.TrimSpace(line)
		switch {
		case isComment(tl) || len(tl) == 0:
			// Do nothing

		case isSectionHeader(tl):
			curSection = make(Section)
			sectionName := strings.TrimSpace(tl[1 : len(tl)-1])
			config[sectionName] = curSection
			curKey = ""
		case isContinuation(line):
			if curKey == "" {
				return nil, fmt.Errorf("Invalid line: %d", lineNum)
			}
			curSection[curKey] = strings.TrimSpace(curSection[curKey] + "\n" + tl)
		default:
			if curSection == nil {
				return nil, fmt.Errorf("Missing section header on line: %d", lineNum)
			}
			// Find closer ':' or ':'
			idx := strings.IndexAny(tl, ":=")

			if idx < 0 {
				return nil, fmt.Errorf("Invalid line: %d", lineNum)
			}

			key, value := strings.TrimSpace(tl[0:idx]), strings.TrimSpace(tl[idx+1:])
			curSection[key] = value
			curKey = key
		}

	}

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		return nil, err
	}

	return
}
