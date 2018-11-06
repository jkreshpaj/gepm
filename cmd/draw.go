package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Box type
type Box []*bytes.Buffer

func getTSize() (int, int) {
	c := exec.Command("stty", "size")
	c.Stdin = os.Stdin
	out, err := c.Output()
	if err != nil {
		fmt.Println(err)
	}
	size := string(out)
	arr := strings.Split(size, " ")
	h, _ := strconv.Atoi(arr[0])
	w := strings.Split(arr[1], "\n")[0]
	wInt, _ := strconv.Atoi(w)
	return wInt, h
}

func setTitle(row *bytes.Buffer) bytes.Buffer {
	_ = "\u001b[1m\u001b[4m\u001b[7m string \u001b[0m"
	name := "Name"
	desc := "Description"
	auth := "Author"
	var c int = 0
	nr := new(bytes.Buffer)
	for _, r := range row.Bytes() {
		if r == 32 {
			c++
		}
	}
	for i := 0; i < c; i++ {
		if i == 0 {
			writeToBuffer(nr, "┃\u001b[1m\u001b[4m\u001b[7m ")
		} else if i == c-1 {
			writeToBuffer(nr, " \u001b[0m┃")
		} else if i < 8 && i > 3 {
			writeToBuffer(nr, string(name[i-4]))
		} else if i > c/3-12 && i < c/3 && c > 90 {
			writeToBuffer(nr, string(desc[i-c/3+11]))
		} else if i > c-27 && i < c-20 {
			writeToBuffer(nr, string(auth[i-c+26]))
		} else {
			writeToBuffer(nr, " ")
		}
	}
	return *nr
}

func writeToBuffer(b *bytes.Buffer, str string) bytes.Buffer {
	_, err := b.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	return *b
}

func writeToRow(b *bytes.Buffer, p Package) bytes.Buffer {
	var c int = 0
	for _, r := range b.Bytes() {
		if r == 32 {
			c++
		}
	}
	for n := 0; n < len(p.Name); n++ {
		b.Bytes()[n+3] = +p.Name[n]
	}
	if len(p.Synopsis) > 0 && c > 90 {
		for d := 0; d < len(p.Synopsis); d++ {
			if d < c-70 {
				b.Bytes()[d+c/3-8] = p.Synopsis[d]
			}
		}
	}
	for d := 0; d < len(p.Author); d++ {
		if d < c-5 {
			b.Bytes()[d+c-23] = p.Author[d]
		}
	}
	return *b
}

func drawBox(w int, h int) Box {
	var box Box
	for i := 0; i < h-5; i++ {
		b := new(bytes.Buffer)
		for j := 0; j < w-1; j++ {
			if j == 0 && i == 0 {
				writeToBuffer(b, "┏")
			} else if i == 0 && j != w-2 {
				writeToBuffer(b, "━")
			} else if i == 0 && j == w-2 {
				writeToBuffer(b, "┓")
			} else if j == w-2 && i != h-6 || j == 0 && i != h-6 {
				writeToBuffer(b, "┃")
			} else if j == 0 && i == h-6 {
				writeToBuffer(b, "┗")
			} else if j == w-2 && i == h-6 {
				writeToBuffer(b, "┛")
			} else if i == h-6 && j != 0 && j != w-2 {
				writeToBuffer(b, "━")
			} else {
				writeToBuffer(b, " ")
			}
		}
		box = append(box, b)
	}
	return box
}
