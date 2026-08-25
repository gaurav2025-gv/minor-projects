package main

import(
    "fmt"
    "io"
    "net/http"
    "os"
    "sync"
)

func download(url string,name string,wg *sync.WaitGroup){

    defer wg.Done()

    resp,err:=http.Get(url)

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    defer resp.Body.Close()

    file,err:=os.Create(name)

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    defer file.Close()

    _,err=io.Copy(file,resp.Body)

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    fmt.Println("Downloaded:",name)
}

func main(){

    downloads:=map[string]string{
        "https://www.example.com/":"file1.html",
        "https://www.google.com/":"file2.html",
        "https://www.github.com/":"file3.html",
    }

    var wg sync.WaitGroup

    for url,name:=range downloads{

        wg.Add(1)

        go download(url,name,&wg)
    }

    wg.Wait()

    fmt.Println("All downloads completed!")
}