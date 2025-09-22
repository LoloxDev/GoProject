package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Hello, qu'est-ce que tu veux ?")
    fmt.Println("1. Bonjour simple")
    fmt.Println("2. Bonjour personnalisé")

    input, _ := reader.ReadString('\n')
    choice, _ := strconv.Atoi(strings.TrimSpace(input))

    switch choice {
    case 1:
        fmt.Println("Hello world!")
    case 2:
        greeter(reader)
    default:
        fmt.Println("Choix invalide")
    }
}

func greeter(reader *bufio.Reader) {
    fmt.Print("Tu t'appelles comment ? ")
    name, _ := reader.ReadString('\n')
    fmt.Printf("Hello, %s\n", strings.TrimSpace(name))
}



