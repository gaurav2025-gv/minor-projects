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

    go receiveMessages(conn)

    reader:=bufio.NewReader(os.Stdin)

    for{

        fmt.Print("Enter message: ")

        message,_:=reader.ReadString('\n')

        conn.Write([]byte(message))
    }
}

func receiveMessages(conn net.Conn){

    buffer:=make([]byte,1024)

    for{

        n,err:=conn.Read(buffer)

        if err!=nil{
            return
        }

        fmt.Println("\nServer:",string(buffer[:n]))
        fmt.Print("Enter message: ")
    }
}