package main

import "fmt"

const englishHelloPrefix = "hello "

func Hello(name string) string {
    if(name != ""){
    
    	return englishHelloPrefix + name
    } else {
        return englishHelloPrefix + "world"
	}
}

func main(){
    fmt.Println(Hello("shm"))
}
