package docs

func init() {

	Definition["set"] = "# `set` - Command Reference\n\n> Define a local variable and set it's value\n\n## Description\n\nDefines, updates or deallocates a local variable.\n\n## Usage\n\n    <stdin> -> set var_name\n    \n    # Assume value from STDIN, define the data type manually\n    <stdin> -> set datatype var_name\n    \n    # Define value manually (data type defaults to string; `str`)\n    set var_name=data\n    \n    # Define value and data type manually\n    set datatype var_name=data\n    \n    # Define a variable but don't set any value\n    set var_name\n    set datatype var_name\n\n## Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> set hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » set hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n## Detail\n\n### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !set var_name\n    \n### Type Annotations\n\nWhen `set` or `global` are used as a function, the parameters are passed as a\nstring which means the variables are defined as a `str`. If you wish to define\nthem as an alternate data type then you should add type annotations:\n\n    » set: int age = 30\n    (`$age` is an integer, `int`)\n\n    » global: bool dark_theme = true\n    \n(`$dark_theme` is a boolean, `bool`)\n\nWhen using `set` or `global` as a method, by default they will define the\nvariable as the data type of the pipe:\n\n    » open: example.json -> set: file\n    \n(`$file` is defined a `json` type because `open` wrote to `set`'s pipe with a\n`json` type)\n\nYou can also annotate `set` and `global` when used as a method too:\n\n    out: 30 -> set: int age\n    \n(`$age` is an integer, `int`, despite `out` writing a string, `str, to the pipe)\n\n> `export` does not support type annotations because environmental variables\n> must always be strings. This is a limitation of the current operating systems.\n\n### Scoping\n\nVariable scoping is simplified to three layers:\n\n1. Local variables (`set`, `!set`, `let`)\n2. Global variables (`global`, `!global`)\n3. Environmental variables (`export`, `!export`, `unset`)\n\nVariables are looked up in that order of too. For example a the following\ncode where `set` overrides both the global and environmental variable:\n\n    » set:    foobar=1\n    » global: foobar=2\n    » export: foobar=3\n    » out: $foobar\n    1\n    \n#### Local variables\n\nThese are defined via `set` and `let`. They're variables that are persistent\nacross any blocks within a function. Functions will typically be blocks\nencapsulated like so:\n\n    function example {\n        # variables scoped inside here\n    }\n    \n...or...\n\n    private example {\n        # variables scoped inside here\n    }\n    \n    \n...however dynamic autocompletes, events, unit tests and any blocks defined in\n`config` will also be triggered as functions.\n\nCode running inside any control flow or error handing structures will be\ntreated as part of the same part of the same scope as the parent function:\n\n    » function example {\n    »     try {\n    »         # set 'foobar' inside a `try` block\n    »         set: foobar=example\n    »     }\n    »     # 'foobar' exists outside of `try` because it is scoped to `function`\n    »     out: $foobar\n    » }\n    example\n    \nWhere this behavior might catch you out is with iteration blocks which create\nvariables, eg `for`, `foreach` and `formap`. Any variables created inside them\nare still shared with any code outside of those structures but still inside the\nfunction block.\n\nAny local variables are only available to that function. If a variable is\ndefined in a parent function that goes on to call child functions, then those\nlocal variables are not inherited but the child functions:\n\n    » function parent {\n    »     # set a local variable\n    »     set: foobar=example\n    »     child\n    » }\n    » \n    » function child {\n    »     # returns the `global` value, \"not set\", because the local `set` isn't inherited\n    »     out: $foobar\n    » }\n    » \n    » global: $foobar=\"not set\"\n    » parent\n    not set\n    \nIt's also worth remembering that any variable defined using `set` in the shells\nFID (ie in the interactive shell) is localised to structures running in the\ninteractive, REPL, shell and are not inherited by any called functions.\n\n#### Global variables\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill be scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\n#### Environmental variables\n\nExported variables (defined via `export`) are system environmental variables.\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> set hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » set cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nThis only works for external executables. There is currently no way to call\naliases, functions nor builtins from a variable and even the above `exec` trick\nis considered bad form because it reduces the readability of your shell scripts.\n\n### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out %($foo)\n    bar\n\n## Synonyms\n\n* `set`\n* `!set`\n\n\n## See Also\n\n* [Reserved Variables](../user-guide/reserved-vars.md):\n  Special variables reserved by _murex_\n* [Variable and Config Scoping](../user-guide/scoping.md):\n  How scoping works within _murex_\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function (deprecated)\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [`expr`](../commands/expr.md):\n  Expressions: mathematical, string comparisons, logical operators\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)"

}
