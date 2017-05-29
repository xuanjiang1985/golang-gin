package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Unix())
}
