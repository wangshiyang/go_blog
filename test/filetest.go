package test

import (
	"bufio"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fi, err := os.Open("E:\\123.txt")

	check(err)

	defer fi.Close()

	read := bufio.NewReader(fi)

	options := make([]string, 1, 1)

	i := 0
	for {
		data, prefix, err := read.ReadLine()

		if err == io.EOF {
			break
		}

		check(err)

		if !prefix {
			//println(string(data))
			options = append(options, string(data))
			i++
		}
	}

	for i = len(options) - 1; i >= 0; i-- {
		println(options[i])
	}
}
