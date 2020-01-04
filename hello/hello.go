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
    return greetingPrefix(language) + name

}
// prefix 是命名返回值，这将在你的函数中创建一个名为 prefix 的变量
func greetingPrefix(language string) (prefix string){

    switch language {
        case french:
            prefix = frenchHelloPrefix
        case spanish:
            prefix = spanishHelloPrefix
        default:
            prefix = englishHelloPrefix
    }
    return   //返回prefix
}

func main(){
    fmt.Println(Hello("shm", ""))
}
