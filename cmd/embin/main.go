package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	defer func(start time.Time) {
		elapsed := time.Since(start)
		log.Printf("Total elapsed time ~ %s", elapsed)
	}(start)

	log.Printf("Starting embin (%v, v%v, build %v)", runtime.Version(), Version(), BuildNumber())

	log.Println("Hello World")
}
