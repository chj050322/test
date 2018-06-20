package main

import "fmt"

var (
	G_BUILD_TIME string
	G_GIT_HASH   string
)

func main() {
	fmt.Printf("%s, %s\n", G_BUILD_TIME, G_GIT_HASH)

}

//修改1
//修改2
//修改3
