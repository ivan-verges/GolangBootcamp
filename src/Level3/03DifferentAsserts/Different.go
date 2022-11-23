package Different

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func StringOk(flag bool) (string, error) {
	if flag {
		return "Ok", nil
	} else {
		return "", errors.New("Empty String")
	}
}
