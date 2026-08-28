package main

import(
    "fmt"
    "time"
)

func scheduleTask(name string,delay time.Duration){

    fmt.Println("Task scheduled:",name)

    time.Sleep(delay)

    fmt.Println("Executing task:",name)
}

func main(){

    fmt.Println("Mini Task Scheduler")

    go scheduleTask("Backup",5*time.Second)
    go scheduleTask("Send Email",3*time.Second)
    go scheduleTask("Generate Report",7*time.Second)

    time.Sleep(8*time.Second)

    fmt.Println("Scheduler finished")
}