package main

import(
    "fmt"
    "os"
)

func getSize(path string)(int64,error){

    entries,err:=os.ReadDir(path)

    if err!=nil{
        return 0,err
    }

    var total int64

    for _,entry:=range entries{

        fullPath:=path+"\\"+entry.Name()

        if entry.IsDir(){

            size,err:=getSize(fullPath)

            if err!=nil{
                return 0,err
            }

            total+=size

        }else{

            info,err:=entry.Info()

            if err!=nil{
                return 0,err
            }

            total+=info.Size()
        }
    }

    return total,nil
}

func main(){

    var path string

    fmt.Print("Enter directory path: ")

    fmt.Scanln(&path)

    size,err:=getSize(path)

    if err!=nil{
        fmt.Println("Error:",err)
        return
    }

    fmt.Printf("Total Size: %.2f MB\n",float64(size)/(1024*1024))
}