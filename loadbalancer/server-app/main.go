package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type serverinfo struct {
    IP              string  `json:"ip"`
    CONTAINER_NAME  string  `json:"container_name"`
}

var si = []serverinfo {
    {IP: "", CONTAINER_NAME: "whatever"},
}

func getServerInfo(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, si)
}

func main() {
    router := gin.Default()
    router.GET("/", getServerInfo)

    router.Run("localhost:3000")
}
