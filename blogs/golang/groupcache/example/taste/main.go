package main

import (
	"fmt"
	"github.com/golang/groupcache"
)

const (
	echoGroupName = "echo"
	echoCacheSize = 1 << 20
)

var echoGroup *groupcache.Group

func main() {
	fmt.Println("main")
	echoGroup = groupcache.NewGroup(echoGroupName, echoCacheSize, groupcache.GetterFunc(func(_ groupcache.Context, key string, dest groupcache.Sink) error {
		return dest.SetString("ECHO:" + key)
	}))

	var b []byte
	err := echoGroup.Get(nil, "hello", groupcache.AllocatingByteSliceSink(&b))
	if err != nil {
		fmt.Errorf("Get error: %s \n", err.Error())
	}

	fmt.Printf("Get hello with:%s \n", string(b))

}
