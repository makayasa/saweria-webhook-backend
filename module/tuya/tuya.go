package tuya

import (
	"fmt"
	"net/http"

	"os/exec"

	"saweria-webhook-backend/env"
	routes "saweria-webhook-backend/routes"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST(routes.Tuya_dimmin, func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		if id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "required ID in param",
			})
			return
		}
		dimMin(id)
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
	r.POST(routes.Tuya_dimmax, func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		if id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "required ID in param",
			})
			return
		}
		dimMax(id)
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
	r.POST(routes.Tuya_turnon, func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		fmt.Println("id:",id)
		if id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "required ID in param",
			})
			return
		}
		turnOn(id)
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
	r.POST(routes.Tuya_turnoff, func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")
		if id == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "required ID in param",
			})
			return
		}
		turnOff(id)
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
}

func dimMin(id string) {
	tuyaInit := ""
	if id == "1" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.set_brightness(10);", env.Lampu_Kamar_ID, env.Lampu_Kamar_IP, env.Lampu_Kamar_Key)
	} else if id == "2" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.set_brightness(10);", env.Lampu_Meja_ID, env.Lampu_Meja_IP, env.Lampu_Meja_Key)
	}
	print("tuya pscript => ", tuyaInit)
	cmd := exec.Command("python", "-c", tuyaInit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(output))
}

func dimMax(id string) {
	tuyaInit := ""
	if id == "1" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.set_brightness(1000);", env.Lampu_Kamar_ID, env.Lampu_Kamar_IP, env.Lampu_Kamar_Key)
	} else if id == "2" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.set_brightness(1000);", env.Lampu_Meja_ID, env.Lampu_Meja_IP, env.Lampu_Meja_Key)
	}
	print("tuya pscript => ", tuyaInit)
	cmd := exec.Command("python", "-c", tuyaInit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(output))
}

func turnOn(id string) {
	tuyaInit := ""
	if id == "1" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.turn_on();", env.Lampu_Kamar_ID, env.Lampu_Kamar_IP, env.Lampu_Kamar_Key)
	} else if id == "2" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.turn_on();", env.Lampu_Meja_ID, env.Lampu_Meja_IP, env.Lampu_Meja_Key)
	}
	cmd := exec.Command("python", "-c", tuyaInit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(output))
}

func turnOff(id string) {
	tuyaInit := ""
	if id == "1" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.turn_off();", env.Lampu_Kamar_ID, env.Lampu_Kamar_IP, env.Lampu_Kamar_Key)
	} else if id == "2" {
		tuyaInit = fmt.Sprintf("import tinytuya; d = tinytuya.BulbDevice('%s', '%s', '%s', version=3.3); d.turn_off();", env.Lampu_Meja_ID, env.Lampu_Meja_IP, env.Lampu_Meja_Key)
	}
	cmd := exec.Command("python", "-c", tuyaInit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(output))
}
