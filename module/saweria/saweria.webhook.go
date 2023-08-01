package saweria

import (
	"fmt"
	routes "saweria-webhook-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/micmonay/keybd_event"
)

func Init(r *gin.Engine) {
	r.POST(routes.Saweria_webhook, func(ctx *gin.Context) {
		body, _ := ctx.GetRawData()
		reqBody := string(body)
		fmt.Println(reqBody)
		// keyboard()
		dropWeapon()
		ctx.JSON(200, gin.H{
			"result": "success",
		})
	})
}

func keyboard() {
	fmt.Println("keyboard masuk")
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	kb.SetKeys(
		keybd_event.VK_A,
	)
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func dropWeapon() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	kb.SetKeys(
		keybd_event.VK_1,
		keybd_event.VK_G,
	)
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}

func allChat() {
}
