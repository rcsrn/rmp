package database

import (
	_"strings"
)

//Search Parser is the parser for specified search by user in the main
//window of the application.
type SearchParser struct {
	text string
}

func CreateNewSearchParser() *SearchParser {
	return &SearchParser{""}
}

func (parser *SearchParser) SetText(text string) {
	parser.text  = text
}

func (parser *SearchParser) Parse() string {
	return ""
}

func splitStatement(request string) []string {
	return nil
}

func obtainTable(word string) string {
	return ""
}

func obtainQueryString(words []string) string {
	return ""
}

func (parser *SearchParser) IsGeneralSearch(text string) bool {
	return string(text[0]) != ":"
}

