package main

import (
    "os"
    "fmt"
    "io"
    "bufio"
    "strings"
)

var ippercentage map[string]int

func main(){
    lf, err := os.OpenFile("log", os.O_RDONLY, 0666)
    if err != nil {
        panic(err)
    }

    defer lf.Close()

    ippercentage = make(map[string]int)
    lineCount := 0
    rd := bufio.NewReader(lf)
    for {
        line, err := rd.ReadString('\n')
        if err != nil && err != io.EOF {
            panic(err)
        }

        if err == io.EOF {
            break
        }

        lineCount += 1
         lns := strings.TrimSpace(line)
        if (lns == "") {
            continue
        }

        ippercentage[lns] = ippercentage[lns] + 1
    }

    for k, v := range ippercentage {
        fmt.Printf("%s - %v %s | count: %v", k, (v*100)/lineCount, "%", v)
        fmt.Print("\n")
    }
}
