package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!if`:             `if`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`and`:             `and`,
	`trypipe`:         `trypipe`,
	`global`:          `global`,
	`murex-docs`:      `murex-docs`,
	`tout`:            `tout`,
	`swivel-table`:    `swivel-table`,
	`>`:               `>`,
	`pt`:              `pt`,
	`if`:              `if`,
	`append`:          `append`,
	`swivel-datatype`: `swivel-datatype`,
	`f`:               `f`,
	`or`:              `or`,
	`try`:             `try`,
	`alter`:           `alter`,
	`brace-quote`:     `brace-quote`,
	`>>`:              `>>`,
	`tread`:           `tread`,
	`export`:          `export`,
	`set`:             `set`,
	`getfile`:         `getfile`,
	`ttyfd`:           `ttyfd`,
	`out`:             `out`,
	`post`:            `post`,
	`err`:             `err`,
	`event`:           `event`,
	`g`:               `g`,
	`rx`:              `rx`,
	`read`:            `read`,
	`catch`:           `catch`,
	`prepend`:         `prepend`,
	`get`:             `get`,
}
