package main

import (
	"log"
	"net/http"
	"os"
    "fmt"

	"github.com/gin-gonic/gin"
)

type serverinfo struct {
	IP             string `json:"ip"`
	CONTAINER_NAME string `json:"container_name"`
    LB_TYPE        string `json:"lb_type"`
}

func getServerInfo(c *gin.Context) {
	var ip string

	f, err := os.OpenFile("/etc/hosts", os.O_RDONLY, 0666)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	b := make([]byte, 1000)
	_, err = f.Read(b)
	if err != nil {
		log.Print(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	ip = string(b[150:162]) // Get the ip inside /etc/hosts file

	if err := f.Close(); err != nil {
		log.Print(err)
	}

	si := serverinfo{
		IP:             ip,
		CONTAINER_NAME: os.Getenv("HOSTNAME"),
        LB_TYPE: os.Getenv("LB_TYPE"),
	}
    log.Print("Response \n %+v", si)
	c.IndentedJSON(http.StatusOK, si)
}

func main() {
	router := gin.Default()
	router.GET("/", getServerInfo)

    var port string

    if port = os.Getenv("PORT"); port == "" {
        port = "3000"
    }
    log.Printf("Current port wil be set as %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
