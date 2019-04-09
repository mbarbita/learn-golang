package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

func execCmd(arg string) []byte {
	var cmdOut []byte
	var err error
	cmd := "curl"
	cmdOut, err = exec.Command(cmd, arg).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return cmdOut
}

func main() {

	// http://user:pass@ip:port/
	arg := "http://192.168.2.12:5984"
	jsonStream := execCmd(arg)
	fmt.Printf("body: %s", jsonStream)

	var x map[string]interface{}
	json.Unmarshal(jsonStream, &x)
	fmt.Printf("UnMarshall: %v\n", x)
	for k, v := range x {
		// fmt.Printf("%s -> %s\n", k, v)
		fmt.Printf("index:%v  value:%v  kind:%s  type:%s\n", k, v, reflect.TypeOf(v).Kind(), reflect.TypeOf(v))
	}
}
