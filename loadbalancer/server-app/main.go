package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type serverinfo struct {
	IP             string `json:"ip"`
	CONTAINER_NAME string `json:"container_name"`
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
	}
	c.IndentedJSON(http.StatusOK, si)
}

func main() {
	router := gin.Default()
	router.GET("/", getServerInfo)

	// listen on 0.0.0.0:3000
	router.Run(":3000")
}
