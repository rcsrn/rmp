package database

import (
	_"strings"
)

//Search Parser is the parser for specified search by user in the main
//window of the application.
//RequestProcessor
type SearchParser struct {
	request string
}

func CreateNewSearchParser(request string) *SearchParser {
	return &SearchParser{request}
}

func (parser *SearchParser) Parse() string {
	words := splitStatement(parser.request)

	query := "SELECT rolas.name FROM"
	
	for i := 0; len(words) > 0; i++ {
		word := words[i]
		if character := word[len(word) - 1]; character == ':' {
			table := obtainTable(word)
			queryString := obtainQueryString(words)

			query += table
		}
	}
	return ""
}

func splitStatement(request string) []string {
	return nil
}

func obtainTable(word string) string {
	return ""
}

func  obtainQueryString(words []string) string {
	return ""
}


func isProcessable(request string) bool {
	return false
}

