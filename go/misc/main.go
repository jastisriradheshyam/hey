package misc

import "strings"

func ParseRunArguments(args string) []string {
	if len(args) == 0 {
		return []string{}
	}
	return strings.Split(args, ".")
}

// Returns module name and task name
func GetModuleAndCommandName(parsedRunArgs []string) (string, string) {
	totalArgParsedLen := len(parsedRunArgs)
	if totalArgParsedLen == 1 {
		return "default", parsedRunArgs[0]
	}

	return strings.Join(parsedRunArgs[0:(totalArgParsedLen-1)], "."), parsedRunArgs[totalArgParsedLen-1]
}

func ParseExportExcludeValues(excludeValues string) []string {
	return []string{}
}
