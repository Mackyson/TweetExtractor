package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	tmp := stdin.Text()
	tmpList := strings.Split(tmp, ",")
	fmt.Println(tmpList)
}
