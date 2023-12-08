package pythonpyautogui

import (
	"fmt"
	"os/exec"
)

func InitPython(chat string) string {
	c := fmt.Sprintf("import pyautogui; pyautogui.write('/all %s');", chat)
	return c

}

func test() {

}

func Execute(c string) {
	cmd := exec.Command("python", "-c", c)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error Init python")
	}
}
