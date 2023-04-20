package docs

func init() {

	Definition["["] = "# `[` (index) - Command Reference\n\n> Outputs an element from an array, map or table\n\n## Description\n\nOutputs an element or multiple elements from an array, map or table.\n\nPlease note that indexes in _murex_ are counted from zero.\n\n## Usage\n\n    <stdin> -> [ element ] -> <stdout>\n    $variable[ element ] -> <stdout>\n    \n    <stdin> -> ![ element ] -> <stdout>\n\n## Examples\n\nReturn the 2nd (1), 4th (3) and 6th (5) element in an array:\n\n    » ja [0..9] -> [ 1 3 5 ]\n    [\n        \"1\",\n        \"3\",\n        \"5\"\n    ]\n    \nReturn the data-type and description of **config shell syntax-highlighting**:\n\n    » config -> [[ /shell/syntax-highlighting ]] -> [ Data-Type Description ]\n    [\n        \"bool\",\n        \"Syntax highlighting of murex code when in the interactive shell\"\n    ]\n    \nReturn all elements _except_ for 1 (2nd), 3 (4th) and 5 (6th):\n\n    » a: [0..9]-> ![ 1 3 5 ]\n    0\n    2\n    4\n    6\n    7\n    8\n    9\n    \nReturn all elements except for the data-type and description:\n\n    » config -> [[ /shell/syntax-highlighting ]] -> ![ Data-Type Description ]\n    {\n        \"Default\": true,\n        \"Dynamic\": false,\n        \"Global\": true,\n        \"Value\": true\n    }\n    \nReturn the top 5 processes from `ps`, ordered by memory usage:\n\n    » ps aux -> [PID %MEM COMMAND] -> sort -nrk2 -> [..5]\n    915961  14.4  /home/lau/dev/go/bin/gopls\n    916184  4.4   /opt/visual-studio-code/code\n    108025  2.9   /usr/lib/firefox/firefox\n    1036    2.4   /usr/lib/baloo_file\n    915710  1.9   /opt/visual-studio-code/code\n    \nReturn the 1st and 30th row:\n\n    » ps aux -> [*1 *30]\n    USER    PID     %CPU    %MEM    VSZ     RSS     TTY     STAT    START   TIME    COMMAND\n    root    37      0.0     0.0     0       0       ?       I<      Dec18   0:00    [kworker/3:0H-events_highpri]\n    \nReturn the 1st and 5th column:\n\n    » ps aux -> [*A *E] -> head -n5                                                                                                                                                                                                       \n    USER    VSZ\n    root    168284\n    root    0\n    root    0\n    root    0\n\n## Detail\n\n### Index counts from zero\n\nIndexes in _murex_ behave like any other computer array in that all arrays\nstart from zero (`0`).\n\n### Include vs exclude\n\nAs demonstrated in the examples above, `[` specifies elements to include\nwhere as `![` specifies elements to exclude.\n\n### Don't error upon missing elements\n\nBy default, **index** generates an error if an element doesn't exist. However\nyou can disable this behavior in `config`\n\n    » config -> [ foobar ]\n    Error in `[` ((builtin) 2,11): Key 'foobar' not found\n    \n    » config set index silent true\n    \n    » config -> [ foobar ]\n\n## Synonyms\n\n* `[`\n* `![`\n* `index`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array"

}
