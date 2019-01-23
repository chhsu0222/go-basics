package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `pasword1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPW := `pasword1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPW))
	if err != nil {
		fmt.Println("You can't login.", err)
		return
	}

	fmt.Println("You are logined.")
}
