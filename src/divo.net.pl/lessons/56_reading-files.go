// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.
package lessons

import(
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadingFiles() {
	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
	
	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	f, err := os.Open("/tmp/dat")
	check(err)
	
	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	
	
}