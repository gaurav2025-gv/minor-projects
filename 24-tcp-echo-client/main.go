package main

import(
    "bufio"
    "fmt"
    "net"
    "os"
)

func main(){

    conn,err:=net.Dial("tcp","localhost:9090")

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    defer conn.Close()

    reader:=bufio.NewReader(os.Stdin)

    for{

        fmt.Print("Enter message: ")

        message,_:=reader.ReadString('\n')

        _,err:=conn.Write([]byte(message))

        if err!=nil{
            fmt.Println("Error:",err)
            return
        }

        buffer:=make([]byte,1024)

        n,err:=conn.Read(buffer)

        if err!=nil{
            fmt.Println("Error:",err)
            return
        }

        fmt.Println("Server:",string(buffer[:n]))
    }
}