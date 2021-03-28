package docs

func init() {

	Definition["function"] = "# _murex_ Shell Docs\n\n## Command Reference: `function`\n\n> Define a function block\n\n## Description\n\n`function` defines a block of code as a function\n\n## Usage\n\n    function: name { code-block }\n    \n    !function: command\n\n## Examples\n\n    » function hw { out \"Hello, World!\" }\n    » hw\n    Hello, World!\n    \n    » !function hw\n    » hw\n    exec: \"hw\": executable file not found in $PATH\n\n## Detail\n\n### Allowed characters\n\nFunction names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n### Undefining a function\n\nLike all other definable states in _murex_, you can delete a function with\nthe bang prefix (see the example above).\n\n### Order of preference\n\nThere is an order of precedence for which commands are looked up:\n1. `test` and `pipe` functions because they alter the behavior of the compiler\n2. Aliases - defined via `alias`. All aliases are global\n3. _murex_ functions - defined via `function`. All functions are global\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n5. variables (dollar prefixed) - declared via `set` or `let`\n6. auto-globbing prefix: `@g`\n7. murex builtins\n8. external executable files\n\n## Synonyms\n\n* `function`\n* `!function`\n\n\n## See Also\n\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [commands/`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [commands/`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [commands/`source` ](../commands/source.md):\n  Import _murex_ code from another file of code block"

}
