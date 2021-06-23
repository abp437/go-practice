package langintro

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cliInput(r *bufio.Reader, p string) (string, error) {
	fmt.Print(p)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// CreateBill creates a bill through user input and then sends it back to the user
func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := cliInput(reader, "Enter your name: ")

	return newBill(name)
}
