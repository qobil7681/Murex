package docs

func init() {

	Definition["runtime"] = "# `runtime` - Command Reference\n\n> Returns runtime information on the internal state of _murex_\n\n## Description\n\n`runtime` is a tool for querying the internal state of _murex_. It's output\nwill be JSON dumps.\n\n## Usage\n\n    runtime: flags -> <stdout>\n    \n`builtins` is an alias for `runtime: --builtins`:\n\n    builtins -> <stdout>\n\n## Examples\n\nList all the builtin data-types that support WriteArray()\n\n    » runtime: --writearray\n    [\n        \"*\",\n        \"commonlog\",\n        \"csexp\",\n        \"hcl\",\n        \"json\",\n        \"jsonl\",\n        \"qs\",\n        \"sexp\",\n        \"str\",\n        \"toml\",\n        \"yaml\"\n    ]\n    \nList all the functions\n\n    » runtime: --functions -> [ agent aliases ]\n    [\n        {\n            \"Block\": \"\\n    # Launch ssh-agent\\n    ssh-agent -\\u003e head -n2 -\\u003e [ :0 ] -\\u003e prefix \\\"export \\\" -\\u003e source\\n    ssh-add: @{g \\u003c!null\\u003e ~/.ssh/*.key} @{g \\u003c!null\\u003e ~/.ssh/*.pem}\\n\",\n            \"FileRef\": {\n                \"Column\": 1,\n                \"Line\": 149,\n                \"Source\": {\n                    \"DateTime\": \"2019-07-07T14:06:11.05581+01:00\",\n                    \"Filename\": \"/home/lau/.murex_profile\",\n                    \"Module\": \"profile/.murex_profile\"\n                }\n            },\n            \"Summary\": \"Launch ssh-agent\"\n        },\n        {\n            \"Block\": \"\\n\\t# Output the aliases in human readable format\\n\\truntime: --aliases -\\u003e formap name alias {\\n        $name -\\u003e sprintf: \\\"%10s =\\u003e ${esccli @alias}\\\\n\\\"\\n\\t} -\\u003e cast str\\n\",\n            \"FileRef\": {\n                \"Column\": 1,\n                \"Line\": 6,\n                \"Source\": {\n                    \"DateTime\": \"2019-07-07T14:06:10.886706796+01:00\",\n                    \"Filename\": \"(builtin)\",\n                    \"Module\": \"source/builtin\"\n                }\n            },\n            \"Summary\": \"Output the aliases in human readable format\"\n        }\n    ]\n    \nTo get a list of every flag supported by `runtime`\n\n    » runtime: --help\n    [\n        \"--aliases\",\n        \"--astcache\",\n        \"--config\",\n        \"--debug\",\n        \"--events\",\n        \"--fids\",\n        \"--flags\",\n        \"--functions\",\n        \"--help\",\n        \"--indexes\",\n        \"--marshallers\",\n        \"--memstats\",\n        \"--modules\",\n        \"--named-pipes\",\n        \"--open-agents\",\n        \"--pipes\",\n        \"--privates\",\n        \"--readarray\",\n        \"--readmap\",\n        \"--sources\",\n        \"--test-results\",\n        \"--tests\",\n        \"--unmarshallers\",\n        \"--variables\",\n        \"--writearray\"\n    ]\n    \nPlease also note that you can supply more than one flag. However when you\ndo use multiple flags the top level of the JSON output will be a map of the\nflag names. eg\n\n    » runtime: --pipes --tests\n    {\n        \"pipes\": [\n            \"file\",\n            \"std\",\n            \"tcp-dial\",\n            \"tcp-listen\",\n            \"udp-dial\",\n            \"udp-listen\"\n        ],\n        \"tests\": {\n            \"state\": {},\n            \"test\": []\n        }\n    }\n    \n    » runtime: --pipes\n    [\n        \"file\",\n        \"std\",\n        \"tcp-dial\",\n        \"tcp-listen\",\n        \"udp-dial\",\n        \"udp-listen\"\n    ]\n    \n    » runtime: --tests\n    {\n        \"state\": {},\n        \"test\": []\n    }\n\n## Flags\n\n* `--aliases`\n    Lists all aliases\n* `--astcache`\n    Lists some data about cached ASTs \n* `--autocomplete`\n    Lists all `autocomplete` schemas - both user defined and automatically generated one\n* `--builtins`\n    Lists all builtin commands, compiled into _murex_\n* `--config`\n    Lists all properties available to `config\n* `--debug`\n    Outputs the state of debug and inspect mode\n* `--events`\n    Lists all builtin event types and any defined events\n* `--exports`\n    Outputs environmental variables. For _murex_ variables (`global` and `set`/`let`) use `--variables\n* `--fids`\n    Lists all running processes / functions\n* `--functions`\n    Lists all _murex_ global functions\n* `--globals`\n    Lists all global variables\n* `--help`\n    Outputs a list of `runtimes`'s flags\n* `--indexes`\n    Lists all builtin data-types which are supported by index (`[`)\n* `--marshallers`\n    Lists all builtin data-types with marshallers (eg required for `format`)\n* `--memstats`\n    Outputs the running state of Go's runtime\n* `--methods`\n    Lists all commands with a defined STDOUT and STDIN data type. This is used to generate smarter autocompletion suggestions with `->\n* `--modules`\n    Lists all installed modules\n* `--named-pipes`\n    Lists all named pipes defined\n* `--not-indexes`\n    Lists all builtin data-types which are supported by index (`![`)\n* `--open-agents`\n    Lists all registered `open` handlers \n* `--pipes`\n    Lists builtin pipes compiled into _murex_. These can be then be defined as named-pipes\n* `--privates`\n    Lists all _murex_ private functions\n* `--readarray`\n    Lists all builtin data-types which support ReadArray()\n* `--readarraywithtype`\n    Lists all builtin data-types which support ReadArrayWithType()\n* `--readmap`\n    Lists all builtin data-types which support ReadMap()\n* `--sources`\n    Lists all loaded murex sources\n* `--summaries`\n    Outputs all the override summaries \n* `--test-results`\n    A dump of any unreported test results\n* `--tests`\n    Lists defined tests\n* `--unmarshallers`\n    Lists all builtin data-types with unmarshallers (eg required for `format`)\n* `--variables`\n    Lists all local _murex_ variables which doesn't include environmental nor global variables\n* `--writearray`\n    Lists all builtin data-types which support WriteArray()\n\n## Detail\n\n### Usage in scripts\n\n`runtime` should not be used in scripts because the output of `runtime` may\nbe subject to change as and when the internal mechanics of _murex_ change.\nThe purpose behind `runtime` is not to provide an API but rather to provide\na verbose \"dump\" of the internal running state of _murex_.\n\nIf you require a stable API to script against then please use the respective\ncommand line tool. For example `fid-list` instead of `runtime --fids`. Some\ntools will provide a human readable output when STDOUT is a TTY but output\na script parsable version when STDOUT is not a terminal.\n\n    » fid-list\n        FID   Parent    Scope  State         Run Mode  BG   Out Pipe    Err Pipe    Command     Parameters\n          0        0        0  Executing     Shell     no                           -murex\n     265499        0        0  Executing     Normal    no   out         err         fid-list\n    \n    » fid-list -> pretty\n    [\n        {\n            \"FID\": 0,\n            \"Parent\": 0,\n            \"Scope\": 0,\n            \"State\": \"Executing\",\n            \"Run Mode\": \"Shell\",\n            \"BG\": false,\n            \"Out Pipe\": \"\",\n            \"Err Pipe\": \"\",\n            \"Command\": \"-murex\",\n            \"Parameters\": \"\"\n        },\n        {\n            \"FID\": 265540,\n            \"Parent\": 0,\n            \"Scope\": 0,\n            \"State\": \"Executing\",\n            \"Run Mode\": \"Normal\",\n            \"BG\": false,\n            \"Out Pipe\": \"out\",\n            \"Err Pipe\": \"err\",\n            \"Command\": \"fid-list\",\n            \"Parameters\": \"\"\n        },\n        {\n            \"FID\": 265541,\n            \"Parent\": 0,\n            \"Scope\": 0,\n            \"State\": \"Executing\",\n            \"Run Mode\": \"Normal\",\n            \"BG\": false,\n            \"Out Pipe\": \"out\",\n            \"Err Pipe\": \"err\",\n            \"Command\": \"pretty\",\n            \"Parameters\": \"\"\n        }\n    ]\n    \n### File reference\n\nSome of the JSON dumps produced from `runtime` will include a map called\n`FileRef`. This is a trace of the source file that defined it. It is used\nby _murex_ to help provide meaningful errors (eg with line and character\npositions) however it is also useful for manually debugging user-defined\nproperties such as which module or script defined an `autocomplete` schema.\n\n### Debug mode\n\nWhen `debug` is enabled garbage collection is disabled for variables and\nFIDs. This means the output of `runtime --variables` and `runtime --fids`\nwill contain more than just the currently defined variables and running\nfunctions.\n\n## Synonyms\n\n* `runtime`\n* `builtins`\n\n\n## See Also\n\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`debug`](../commands/debug.md):\n  Debugging information\n* [`event`](../commands/event.md):\n  Event driven programming for shell scripts\n* [`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [`fid-list`](../commands/fid-list.md):\n  Lists all running functions within the current _murex_ session\n* [`foreach`](../commands/foreach.md):\n  Iterate through an array\n* [`formap`](../commands/formap.md):\n  Iterate through a map or other collection of data\n* [`format`](../commands/format.md):\n  Reformat one data-type into another data-type\n* [`function`](../commands/function.md):\n  Define a function block\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`method`](../commands/method.md):\n  Define a methods supported data-types\n* [`open`](../commands/open.md):\n  Open a file with a preferred handler\n* [`openagent`](../commands/openagent.md):\n  Creates a handler function for `open\n* [`pipe`](../commands/pipe.md):\n  Manage _murex_ named pipes\n* [`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block\n* [`test`](../commands/test.md):\n  _murex_'s test framework - define tests, run tests and debug shell scripts"

}
