package main

import(
    "encoding/json"
    "fmt"
    "os"
)

type Config struct{
    AppName string `json:"app_name"`
    Port int `json:"port"`
    Debug bool `json:"debug"`
}

func main(){

    data,err:=os.ReadFile("config.json")

    if err!=nil{
        fmt.Println("Error reading config:",err)
        return
    }

    var config Config

    err=json.Unmarshal(data,&config)

    if err!=nil{
        fmt.Println("Error parsing JSON:",err)
        return
    }

    fmt.Println("App Name:",config.AppName)
    fmt.Println("Port:",config.Port)
    fmt.Println("Debug:",config.Debug)

    config.Port=9090

    newData,err:=json.MarshalIndent(config,"","    ")

    if err!=nil{
        fmt.Println("Error creating JSON:",err)
        return
    }

    err=os.WriteFile("config.json",newData,0644)

    if err!=nil{
        fmt.Println("Error writing config:",err)
        return
    }

    fmt.Println("Config updated successfully!")
}