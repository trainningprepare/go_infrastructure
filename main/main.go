package main

import (
    consistent "go_infrastructure/hash"

	"fmt"
)

func main() {

	ring:=consistent.CreateConsitentRing(100)

	keynames:=make([]string,100)

	for key, _ := range keynames {
		keynames[key] = fmt.Sprintf("nodename_%d",key)
	}

	for i := 0; i < 100; i++ {
		ring.AddBucket(keynames[i])
	}

	ring.DumpNodesRange()



}