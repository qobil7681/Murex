package docs

func init() {

	Definition["prefix"] = "# _murex_ Shell Docs\n\n## Command Reference: `prefix`\n\n> Prefix a string to every item in a list\n\n## Description\n\nTakes a list from STDIN and returns that same list with each element prefixed.\n\n## Usage\n\n    <stdin> -> prefix str -> <stdout>\n\n## Examples\n\n    » ja: [Monday..Wednesday] -> prefix foobar\n    [\n        \"foobarMonday\",\n        \"foobarTuesday\",\n        \"foobarWednesday\"\n    ]\n\n## Detail\n\nSupported data types can queried via `runtime`\n\n    runtime: --marshallers\n    runtime: --unmarshallers\n\n## See Also\n\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [apis/`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):\n  Converts structured memory into a _murex_ data-type (eg for stdio)\n* [apis/`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):\n  Converts a _murex_ data-type into structured memory\n* [commands/`left`](../commands/left.md):\n  Left substring every item in a list\n* [commands/`len` ](../commands/len.md):\n  Outputs the length of an array\n* [commands/`right`](../commands/right.md):\n  Right substring every item in a list\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [commands/`suffix`](../commands/suffix.md):\n  Prefix a string to every item in a list"

}
