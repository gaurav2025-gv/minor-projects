package main

import(
	"bufio"
	"fmt"
	"net"
)
func main() {
	conn,err:=net.Dial("tcp","localhost:9090")
	if(err!=nil) {
		fmt.Println("Error",err)
		return
	}
	defer conn.Close()

	request:="GET / HTTP/1.1\r\nHost: localhost\r\nConnection: close\r\n\r\n"
	  _,err=conn.Write([]byte(request))
    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    reader:=bufio.NewReader(conn)

    for{

        line,err:=reader.ReadString('\n')

        if err!=nil{
            break
        }

        fmt.Print(line)
    }
}

