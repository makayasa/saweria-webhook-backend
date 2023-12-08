package saweria

import (
	"fmt"
	routes "saweria-webhook-backend/routes"
	"time"

	// pythonpyautogui "saweria-webhook-backend/utils/python-pyautogui"
	stringtokeyboardgo "saweria-webhook-backend/utils/stringtokeyboard.go"

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
		// allChat()
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
	)
	kb.Launching()
	time.Sleep(20 * time.Millisecond)
	kb.SetKeys(
		keybd_event.VK_G,
	)
	kb.Launching()
}

func allChat() {
	fmt.Println("msk all-chat")
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	kb.SetKeys(
		keybd_event.VK_ENTER,
	)
	kb.Launching()

	// stringtokeyboardgo.KeyboardWrite("/all")
	stringtokeyboardgo.KeyboardWrite("asadasan")
	// chat := "chamber kontrol"
	// c := pythonpyautogui.InitPython(chat)
	// pythonpyautogui.Execute(c)
	kb.HasSHIFT(true)
	kb.HasSHIFTR(true)
	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Launching()
	time.Sleep(20 * time.Millisecond)
	kb.HasSHIFT(false)
	kb.HasSHIFTR(false)
	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Launching()
}

// /all asadasan
