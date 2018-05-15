package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!global`:         `global`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!event`:          `event`,
	`append`:          `append`,
	`rx`:              `rx`,
	`set`:             `set`,
	`export`:          `export`,
	`tout`:            `tout`,
	`and`:             `and`,
	`try`:             `try`,
	`trypipe`:         `trypipe`,
	`murex-docs`:      `murex-docs`,
	`brace-quote`:     `brace-quote`,
	`g`:               `g`,
	`read`:            `read`,
	`prepend`:         `prepend`,
	`post`:            `post`,
	`pt`:              `pt`,
	`event`:           `event`,
	`get`:             `get`,
	`out`:             `out`,
	`ttyfd`:           `ttyfd`,
	`catch`:           `catch`,
	`swivel-table`:    `swivel-table`,
	`getfile`:         `getfile`,
	`>`:               `>`,
	`tread`:           `tread`,
	`f`:               `f`,
	`or`:              `or`,
	`if`:              `if`,
	`global`:          `global`,
	`alter`:           `alter`,
	`swivel-datatype`: `swivel-datatype`,
	`err`:             `err`,
	`>>`:              `>>`,
}
