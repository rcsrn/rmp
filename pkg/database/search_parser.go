package database

//RequestProcessor is the parser for specified search by user in the main
//window of the application.
//RequestProcessor
type SearchParser struct {
	statement string
}

func CreateNewSearchParser(statement string) *SearchParser {
	return &SearchParser{statement}
}

func Parse(request string) string {
	return ""
}

func isProcessable(request string) bool {
	return false
}
