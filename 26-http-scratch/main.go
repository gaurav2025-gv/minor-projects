package main

import(
	"bufio"
	"fmt"
	"net"
	"strings"
)
func main() {
	listener,err:=net.Listen("tcp",":9090")
	if err!=nil{
	fmt.Println("Error",err)
	return
 }
 defer listener.Close()
 fmt.Println("HTTP Server running on port 9090")
  for{

        conn,err:=listener.Accept()

        if err!=nil{
            fmt.Println("Error:",err)
            continue
        }

        go handleConnection(conn)
    }
}
func handleConnection(conn net.Conn){

    defer conn.Close()

    reader:=bufio.NewReader(conn)

    requestLine,err:=reader.ReadString('\n')

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    requestLine=strings.TrimSpace(requestLine)

    parts:=strings.Fields(requestLine)

    if len(parts)<3{
        fmt.Println("Malformed request")
        return
    }

    method:=parts[0]
    path:=parts[1]
    version:=parts[2]

    fmt.Println("Method:",method)
    fmt.Println("Path:",path)
    fmt.Println("Version:",version)

    body:="<h1>Hello, World!</h1><p>Ye mera pehla HTTP server hai</p>"

response:=fmt.Sprintf(
    "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: %d\r\n\r\n%s",
    len(body),
    body,
)

    conn.Write([]byte(response))
}
