package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Encrypt (Enter 1)? Decrypt (Enter 2): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.ParseInt(input, 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	if num == 1 {
		fmt.Println("You chose Encrypt!")
	} else if num == 2 {
		fmt.Println("You chose Decrypt!")
	} else {
		fmt.Println("Incorrect choice!")
	}
	fmt.Println(input)
}
