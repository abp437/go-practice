package langintro

import (
	"fmt"
	"io/ioutil"
)

func writeToFile(s string, fn string) {
	data := []byte(s)

	err := ioutil.WriteFile("./tmp/"+fn+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Saved to file")
}
