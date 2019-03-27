package main

import (
	"path/filepath"
	"strings"
)

// remove quotes, single quotes, and semicolons
var replacer = strings.NewReplacer("\"", "", "'", "", ";", "")

func starts(search string, line string) bool {
	if len(line) < len(search) {
		return false
	}
	return line[:len(search)] == search
}

func after(search string, line string) string {
	if len(line) < len(search) {
		return ""
	}
	line = line[len(search):]
	line = replacer.Replace(line)
	return filepath.Clean(line)
}

func trueFile(callFrom string, importName string) string {
	importName = replacer.Replace(importName)
	pos := strings.LastIndex(callFrom, "/")
	ret := callFrom[:pos+1] + importName
	return filepath.Clean(ret)
}
