package main

import(
    "fmt"
    "net"
)

func main(){

    listener,err:=net.Listen("tcp",":9090")

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    defer listener.Close()

    fmt.Println("TCP Echo Server running on port 9090...")

    for{

        conn,err:=listener.Accept()

        if err!=nil{
            fmt.Println("Error:",err)
            continue
        }

        fmt.Println("Client connected")

        handleConnection(conn)
    }
}

func handleConnection(conn net.Conn){

    defer conn.Close()

    buffer:=make([]byte,1024)

    for{

        n,err:=conn.Read(buffer)

        if err!=nil{
            return
        }

        message:=string(buffer[:n])

        fmt.Println("Received:",message)

        conn.Write([]byte(message))
    }
}