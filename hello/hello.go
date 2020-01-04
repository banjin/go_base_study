package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const englishHelloPrefix = "hello "

func Hello(name string, language string) string {
    if name == ""{
        name = "world"
    }

    prefix := englishHelloPrefix
    switch language {
        case french:
            prefix = frenchHelloPrefix
        case spanish:
            prefix = spanishHelloPrefix
    }
    return prefix + name
}

func main(){
    fmt.Println(Hello("shm", ""))
}
