package arg

import (
	"os"
	"regexp"
	"strings"
)

var gArgs []string
var gBlArgs = make(map[string]bool)
var gKvArgs = make(map[string]string)

func Args() []string {
	return gArgs
}

func BoolArgs() map[string]bool {
	return gBlArgs
}

func KeyValueArgs() map[string]string {
	return gKvArgs
}

func Parse() {
	if len(os.Args) < 2 {
		return
	}
	rawArgs := os.Args[1:]
	wordRegex := regexp.MustCompile("^[a-z]+$")
	for _, arg := range rawArgs {
		if strings.HasPrefix(arg, "--") {
			rest := arg[2:]
			if wordRegex.MatchString(rest) {
				gBlArgs[rest] = true
				continue
			} else if items := strings.SplitN(rest, "=", 2); len(items) == 2 && wordRegex.MatchString(items[0]) {
				gKvArgs[items[0]] = items[1]
				continue
			}
		}
		gArgs = append(gArgs, arg)
	}
}
