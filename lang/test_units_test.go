package lang_test

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

var uniq int32

type testUTPs struct {
	Function  string
	TestBlock string
	Passed    bool
	UTP       lang.UnitTestPlan
}

func testRunTest(t *testing.T, plans []testUTPs) {
	count.Tests(t, len(plans)*2, "testRunTest")

	lang.InitEnv()

	var pubpriv string

	for i := range plans {
		for j := 1; j < 3; j++ { // test public functions the private functions

			fileRef := &ref.File{
				Source: &ref.Source{
					Filename: "foobar.mx",
					Module:   "foobar/mod-" + strconv.Itoa(int(atomic.AddInt32(&uniq, 1))),
					DateTime: time.Now(),
				},
			}

			if j == 1 {
				lang.MxFunctions.Define(plans[i].Function, []rune(plans[i].TestBlock), fileRef)
				pubpriv = "public"
			} else {
				lang.PrivateFunctions.Define(plans[i].Function, []rune(plans[i].TestBlock), fileRef)
				plans[i].Function = fileRef.Source.Module + "/" + plans[i].Function
				pubpriv = "private"
			}

			ut := new(lang.UnitTests)
			ut.Add(plans[i].Function, &plans[i].UTP, fileRef)

			if ut.Run(lang.ShellProcess.Tests, plans[i].Function) != plans[i].Passed {
				if plans[i].Passed {
					t.Errorf("Unit test %s %d failed", pubpriv, i)
					b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
					if err != nil {
						panic(err)
					}

					t.Logf("Test report:\n%s", b)

				} else {
					t.Errorf("Unit test %s %d passed when expected to fail", pubpriv, i)
				}
			}
			lang.ShellProcess.Tests.Results = new(lang.TestResults)
		}
	}
}
