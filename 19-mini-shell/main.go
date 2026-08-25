package main

import(
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main(){

    reader:=bufio.NewReader(os.Stdin)

    fmt.Println("Mini Shell Started")

    for{

        fmt.Print("myshell> ")

        input,_:=reader.ReadString('\n')

        input=strings.TrimSpace(input)

        if input=="exit"{
            fmt.Println("Goodbye!")
            break
        }

        if input==""{
            continue
        }

        command:=exec.Command("cmd","/C",input)

        command.Stdout=os.Stdout
        command.Stderr=os.Stderr

        err:=command.Run()

        if err!=nil{
            fmt.Println("Error:",err)
        }
    }
}