package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`echo`:            `out`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!event`:          `event`,
	`if`:              `if`,
	`append`:          `append`,
	`getfile`:         `getfile`,
	`out`:             `out`,
	`f`:               `f`,
	`rx`:              `rx`,
	`global`:          `global`,
	`get`:             `get`,
	`post`:            `post`,
	`brace-quote`:     `brace-quote`,
	`g`:               `g`,
	`set`:             `set`,
	`swivel-table`:    `swivel-table`,
	`ttyfd`:           `ttyfd`,
	`catch`:           `catch`,
	`export`:          `export`,
	`err`:             `err`,
	`read`:            `read`,
	`trypipe`:         `trypipe`,
	`alter`:           `alter`,
	`prepend`:         `prepend`,
	`murex-docs`:      `murex-docs`,
	`>>`:              `>>`,
	`swivel-datatype`: `swivel-datatype`,
	`tout`:            `tout`,
	`try`:             `try`,
	`or`:              `or`,
	`event`:           `event`,
	`>`:               `>`,
	`pt`:              `pt`,
	`tread`:           `tread`,
	`and`:             `and`,
}
