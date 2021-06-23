package langintro

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CreateBill creates a bill through user input and then sends it back to the user
func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return newBill(name)
}
