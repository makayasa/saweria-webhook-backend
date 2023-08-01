package main

import (
	"fmt"
	"os/exec"
	conts "saweria-webhook-backend/const"
	saweria "saweria-webhook-backend/module/saweria"
	tuya "saweria-webhook-backend/module/tuya"

	"github.com/gin-gonic/gin"
)

func main() {
	runNgrok()
	var r = gin.Default()
	tuya.Init(r)
	saweria.Init(r)
	address := fmt.Sprintf("%s:%d", conts.Hostname, conts.Port)
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "nice",
		})
	})
	r.Run(address)
	return
}

func runNgrok() {
	fmt.Println("runNgrok")
	ngrok := exec.Command("ngrok", "http", "--domain=turkey-amazed-regularly.ngrok-free.app", "192.168.0.2:7070")
	go ngrok.CombinedOutput()

}
