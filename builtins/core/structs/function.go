package structs

import (
	"errors"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["alias"] = cmdAlias
	lang.GoFunctions["!alias"] = cmdUnalias
	lang.GoFunctions["func"] = cmdFunc
	lang.GoFunctions["!func"] = cmdUnfunc
	lang.GoFunctions["private"] = cmdPrivate
	lang.GoFunctions["!private"] = cmdUnprivate
}

var rxAlias = regexp.MustCompile(`^([_a-zA-Z0-9]+)=(.*?)$`)

func cmdAlias(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)
		b, err := json.Marshal(lang.GlobalAliases.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		return err

	}

	p.Stdout.SetDataType(types.Null)

	s, _ := p.Parameters.String(0)

	if !rxAlias.MatchString(s) {
		return errors.New("Invalid syntax. Expecting `alias new_name=original_name parameter1 parameter2 ...`")
	}

	split := rxAlias.FindStringSubmatch(s)
	name := split[1]
	params := append([]string{split[2]}, p.Parameters.StringArray()[1:]...)

	lang.GlobalAliases.Add(name, params)
	return nil
}

func cmdUnalias(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for _, name := range p.Parameters.StringArray() {
		err := lang.GlobalAliases.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmdFunc(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	lang.MxFunctions.Define(name, p.Module, block)
	return nil
}

func cmdUnfunc(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return lang.MxFunctions.Undefine(name)
}

func cmdPrivate(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	lang.PrivateFunctions.Define(name, p.Module, block)
	return nil
}

func cmdUnprivate(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return lang.PrivateFunctions.Undefine(name)
}
