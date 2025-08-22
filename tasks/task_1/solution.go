package task_1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TaskOne() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fs := bufio.NewReadWriter(in, out)
	defer fs.Flush()

	countOfHexagons := 0
	fmt.Fscan(fs, &countOfHexagons)

	for i := 0; i < countOfHexagons; i++ {
		printHexagon(fs)
	}
}

func printHexagon(fs *bufio.ReadWriter) {
	width, height := 0, 0
	fmt.Fscan(fs, &width, &height)

	strLine := strings.Builder{}
	fullHeight := height*2 + 1
	step := 0

	for row := 0; row < fullHeight; row++ {
		strLine.Reset()
		if row == 0 {
			strLine.WriteString(strings.Repeat(" ", height))
			strLine.WriteString(strings.Repeat("_", width))
			step++
		} else if row <= height {
			strLine.WriteString(strings.Repeat(" ", height-step))
			strLine.WriteByte('/')
			strLine.WriteString(strings.Repeat(" ", width+(step-1)*2))
			strLine.WriteByte('\\')
			step++
		} else {
			step--
			strLine.WriteString(strings.Repeat(" ", height-step))
			strLine.WriteByte('\\')
			if step == 1 {
				strLine.WriteString(strings.Repeat("_", width))
			} else {
				strLine.WriteString(strings.Repeat(" ", width+(step-1)*2))
			}
			strLine.WriteByte('/')
		}
		strLine.WriteByte('\n')
		fmt.Fprint(fs, strLine.String())
	}
}
