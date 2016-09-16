// DrawRectangle
package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// create the local error and file variables
	var file *os.File
	var err error
	filename := "pic.txt"

	// create the file to dump the info and set permissions
	file, err = os.Create(filename)
	check(err)
	file.Write([]byte{'0'})
	check(file.Close())

	for i := 0; i < 10; i++ {
		// re-open the file and check for errors
		file, err = os.OpenFile(filename, syscall.O_WRONLY, 0777)
		check(err)

		// create rectangle
		bytes := []byte{byte(i)}
		_, err = file.Write(bytes)
		check(err)
		check(file.Sync())
		check(file.Close())

		// notify user and wait for a half a second
		fmt.Println("wrote png " + strconv.Itoa(int(i)))
		time.Sleep(500 * time.Millisecond)
	}
}
