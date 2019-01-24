package docs

func init() {
	Summary = map[string]string{

		"!":                     "Reads the STDIN and exit number from previous process and not's it's condition",
		"(":                     "Write a string to the STDOUT without new line",
		"2darray":               "Create a 2D JSON array from multiple input sources",
		">>":                    "Writes STDIN to disk - appending contents if file already exists",
		">":                     "Writes STDIN to disk - overwriting contents if file already exists",
		"@[":                    "Outputs a ranged subset of data from STDIN",
		"[":                     "Outputs an element from an array, map or table",
		"a":                     "A sophisticated yet simply way to build an array or list",
		"alter":                 "Change a value within a structured data-type and pass that change along the pipeline without altering the original source input",
		"and":                   "Returns `true` or `false` depending on whether multiple conditions are met",
		"append":                "Add data to the end of an array",
		"cast":                  "Alters the data type of the previous function without altering it's output",
		"catch":                 "Handles the exception code raised by `try` or `trypipe",
		"cd":                    "Change (working) directory",
		"cpuarch":               "Output the hosts CPU architecture",
		"cpucount":              "Output the number of CPU cores available on your host",
		"debug":                 "Debugging information",
		"die":                   "Terminate murex with an exit number of 1",
		"err":                   "Print a line to the STDERR",
		"exec":                  "Runs an executable",
		"exit":                  "Exit murex",
		"exitnum":               "Output the exit number of the previous process",
		"export":                "Define a local variable and set it's value",
		"f":                     "Lists objects (eg files) in the current working directory",
		"false":                 "Returns a `false` value",
		"g":                     "Glob pattern matching for file system objects (eg *.txt)",
		"get":                   "Makes a standard HTTP request and returns the result as a JSON object",
		"getfile":               "Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.",
		"global":                "Define a global variable and set it's value",
		"history":               "Outputs murex's command history",
		"if":                    "Conditional statement to execute different blocks of code depending on the result of the condition",
		"ja":                    "A sophisticated yet simply way to build a JSON array",
		"jsplit":                "Splits STDIN into a JSON array based on a regex parameter",
		"len":                   "Outputs the length of an array",
		"lockfile":              "Create and manage lock files",
		"man-summary":           "Outputs a man page summary of a command",
		"map":                   "Creates a map from two data sources",
		"murex-docs":            "Displays the man pages for _murex_ builtins",
		"murex-update-exe-list": "Forces _murex_ to rescan $PATH looking for exectables",
		"null":                  "null function. Similar to /dev/null",
		"or":                    "Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.",
		"os":                    "Output the auto-detected OS name",
		"out":                   "`echo` a string to the STDOUT with a trailing new line character",
		"post":                  "HTTP POST request with a JSON-parsable return",
		"prepend":               "Add data to the start of an array",
		"pretty":                "Prettifies JSON to make it human readable",
		"pt":                    "Pipe telemetry. Writes data-types and bytes written",
		"read":                  "`read` a line of input from the user and store as a variable",
		"rx":                    "Regexp pattern matching for file system objects (eg '.*\\.txt')",
		"set":                   "Define a local variable and set it's value",
		"swivel-datatype":       "Converts tabulated data into a map of values for serialised data-types such as JSON and YAML",
		"swivel-table":          "Rotates a table by 90 degrees",
		"tout":                  "Print a string to the STDOUT and set it's data-type",
		"tread":                 "`read` a line of input from the user and store as a user defined *typed* variable",
		"true":                  "Returns a `true` value",
		"try":                   "Handles errors inside a block of code",
		"trypipe":               "Checks state of each function in a pipeline and exits block on error",
	}

	Synonym = map[string]string{

		"!":                     "!",
		"(":                     "(",
		"2darray":               "2darray",
		">>":                    ">>",
		">":                     ">",
		"@[":                    "@[",
		"[":                     "[",
		"![":                    "[",
		"a":                     "a",
		"alter":                 "alter",
		"and":                   "and",
		"!and":                  "and",
		"append":                "append",
		"cast":                  "cast",
		"catch":                 "catch",
		"!catch":                "catch",
		"cd":                    "cd",
		"cpuarch":               "cpuarch",
		"cpucount":              "cpucount",
		"debug":                 "debug",
		"die":                   "die",
		"err":                   "err",
		"exec":                  "exec",
		"exit":                  "exit",
		"exitnum":               "exitnum",
		"export":                "export",
		"!export":               "export",
		"unset":                 "export",
		"f":                     "f",
		"false":                 "false",
		"g":                     "g",
		"get":                   "get",
		"getfile":               "getfile",
		"global":                "global",
		"!global":               "global",
		"history":               "history",
		"if":                    "if",
		"!if":                   "if",
		"ja":                    "ja",
		"jsplit":                "jsplit",
		"len":                   "len",
		"lockfile":              "lockfile",
		"man-summary":           "man-summary",
		"map":                   "map",
		"murex-docs":            "murex-docs",
		"murex-update-exe-list": "murex-update-exe-list",
		"null":                  "null",
		"or":                    "or",
		"!or":                   "or",
		"os":                    "os",
		"out":                   "out",
		"echo":                  "out",
		"post":                  "post",
		"prepend":               "prepend",
		"pretty":                "pretty",
		"pt":                    "pt",
		"read":                  "read",
		"rx":                    "rx",
		"set":                   "set",
		"!set":                  "set",
		"swivel-datatype":       "swivel-datatype",
		"swivel-table":          "swivel-table",
		"tout":                  "tout",
		"tread":                 "tread",
		"true":                  "true",
		"try":                   "try",
		"trypipe":               "trypipe",
	}
}
