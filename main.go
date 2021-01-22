package main

import (
	"log"

	"commonDemo/lib"
)

func main() {
	if err := lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"}); err != nil {
		log.Fatal(err)
	}
	defer lib.Destroy()

	// todo sth
	// lib.Log.TagInfo(lib.NewTrace(), lib.DLTagUndefind, map[])
}
