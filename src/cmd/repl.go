package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	for {
		v := rep()
		s := strings.TrimSpace(string(v))
		if s == "exit"{
			break
		}
		exec([]byte(s))
	}
}

func trim(v string) string{
	var tr  []byte = nil
	t := []byte(v)
	for i := 0; i < len(v); i++{
		if t[i] == ' ' {
			tr = t[i+1:]
		}else {
			break
		}
	}
	return string(tr)
}

func exec(v []byte){
	fmt.Println("")
	fmt.Print(string(v))

}
func rep() []byte{
	in := os.Stdin
	fmt.Print(">")
	s := make([]byte,1)
	c,err := in.Read(s)
	buffer := make([]byte,1)
	for ; err == nil; c,err = in.Read(s){
		if c == 0 {
			break
		}
		if s[0] == 0 {
			continue
		}
		if s[0] == '\r' || s[0] == '\n' {
			break
		}
		print(s[0])
		buffer = append(buffer,s[0])
		s = make([]byte,1)
	}
	in.Sync()
	return buffer
}
