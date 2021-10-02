package datatools

import (
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

func init() {
	lang.DefineMethod("alter", cmdAlter, types.Unmarshal, types.Marshal)

	defaults.AppendProfile(`
		autocomplete: set alter { [{
			"AnyValue": true,
			"ExecCmdline": true,
			"AutoBranch": true,
			"Dynamic": ({ -> struct-keys }),
			"Flags": [ "-m", "--merge" ],
			"FlagValues": {
				"-m": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}],
				"--merge": [{
					"AnyValue": true,
					"ExecCmdline": true,
					"AutoBranch": true,
					"Dynamic": ({ -> struct-keys })
				}]
			}
		} ]}
	`)
}

func cmdAlter(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	var (
		merge  bool
		offset int
	)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if s == "-m" || s == "--merge" {
		merge = true
		offset++

		s, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
	}

	new, err := p.Parameters.String(1 + offset)
	if err != nil {
		return err
	}

	path, err := alter.SplitPath(s)
	if err != nil {
		return err
	}

	if merge {
		v, err = alter.Merge(p.Context, v, path, new)
		if err != nil {
			return err
		}
	} else {
		v, err = alter.Alter(p.Context, v, path, new)
		if err != nil {
			return err
		}
	}

	b, err := lang.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
