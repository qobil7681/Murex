package generic

import (
	"bufio"
	"bytes"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
)

func readArray(read stdio.Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(bytes.TrimSpace(scanner.Bytes()))
	}

	err := scanner.Err()
	return err
}
