//// +build openbsd

package man

import (
	"bufio"
	"compress/gzip"
	"github.com/lmorg/murex/utils"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	rxMatchFlagsEscaped *regexp.Regexp = regexp.MustCompile(`\\f[BI](\\-[a-zA-Z0-9]|\\-\\-[\\\-a-zA-Z0-9]+).*?\\f[RP]`)
	rxMatchFlagsQuoted  *regexp.Regexp = regexp.MustCompile(`\.IP "(.*?)"`)
	rxMatchFlagsFlag    *regexp.Regexp = regexp.MustCompile(`(--[\-a-zA-Z0-9]+)`)
)

/*
	1	General commands (tools and utilities).
	2	System calls and error numbers.
	3	Library functions.
	3p	perl(1) programmer's reference guide.
	4	Device drivers.
	5	File formats.
	6	Games.
	7	Miscellaneous information.
	8	System maintenance and operation commands.
	9	Kernel internals.
*/

func Initialise() {}

func ScanManPages(exe string) (flags []string) {
	// Get paths
	cmd := exec.Command("man", "-w", exe)
	b, err := cmd.Output()
	if err != nil {
		os.Stderr.WriteString("Error initialising man pages: " + err.Error() + utils.NewLineString)
		return
	}

	s := strings.TrimSpace(string(b))
	paths := strings.Split(s, ":")

	// Parse man pages
	fMap := make(map[string]bool)
	for i := range paths {
		parseManPage(&fMap, paths[i])
	}

	for f := range fMap {
		flags = append(flags, f)
	}
	sort.Strings(flags)
	return
}

func parseManPage(flags *map[string]bool, filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return
	}

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		defer gz.Close()
		if err != nil {
			return
		}

		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		s := scanner.Text()

		match := rxMatchFlagsEscaped.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			s := strings.Replace(match[i][1], `\`, "", -1)
			if strings.HasSuffix(s, "fR") || strings.HasSuffix(s, "fP") {
				s = s[:len(s)-2]
			}
			(*flags)[s] = true
		}

		match = rxMatchFlagsQuoted.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			flag := rxMatchFlagsFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = true
			}
		}
	}

	return
}
