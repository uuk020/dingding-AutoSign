package main

import (
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

func AdbShellDumpsysPowerOff() bool {
	flag := false
	MyCmd := exec.Command("cmd.exe", "/c", "adb shell dumpsys power")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	if strings.Contains(s, "Display Power: state=OFF") {
		flag = true
	}
	return flag
}

func AdbShellInputTap(x, y int) {
	x2 := strconv.Itoa(x)
	y2 := strconv.Itoa(y)
	exec.Command("adb", "shell", "input", "tap", x2, y2).Run()
}

func AdbShellInputSwipe(x1, y1, x2, y2 int) {
	xx1 := strconv.Itoa(x1)
	yy1 := strconv.Itoa(y1)
	xx2 := strconv.Itoa(x2)
	yy2 := strconv.Itoa(y2)
	exec.Command("adb", "shell", "input", "swipe", xx1, yy1, xx2, yy2).Run()
}

func AdbShellInputKeyEvent(s string) {
	exec.Command("adb", "shell", "input", "keyevent", s).Run()
}

func AdbShellInputText(s string) {
	exec.Command("adb", "shell", "input", "text", s).Run()
}
