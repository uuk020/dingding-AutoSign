package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// config 结构体
type Config struct {
	Lock          Lock          `json:"lock"`
	StartDingding StartDingding `json:"start_dingding"`
	WorkPlatform  WorkPlatform  `json:"work_platform"`
	Sign          Sign          `json:"sign"`
}
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Lock struct {
	Require    bool       `json:"require"`
	Password   string     `json:"password"`
	Coordinate Coordinate `json:"coordinate"`
}
type StartDingding struct {
	Require    bool       `json:"require"`
	Coordinate Coordinate `json:"coordinate"`
}
type WorkPlatform struct {
	Coordinate Coordinate `json:"coordinate"`
}
type Click struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Photo struct {
	Require    bool       `json:"require"`
	Coordinate Coordinate `json:"coordinate"`
}
type Sign struct {
	Coordinate Coordinate `json:"coordinate"`
	Click      Click      `json:"click"`
	Photo      Photo      `json:"photo"`
}

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

func main() {
	fmt.Println("开启钉钉打卡~")
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal([]byte(byteValue), &config)

	// 检查是否熄屏
	if AdbShellDumpsysPowerOff() {
		AdbShellInputKeyEvent("26") //power
	}
	// 滑动解锁
	AdbShellInputSwipe(747, 1413, 747, 323)
	if config.Lock.Require {
		AdbShellInputText(config.Lock.Password)
		AdbShellInputTap(config.Lock.Coordinate.X, config.Lock.Coordinate.Y)
	}
	// 开启钉钉，若在钉钉界面则不需要
	if config.StartDingding.Require {
		AdbShellInputTap(config.StartDingding.Coordinate.X, config.StartDingding.Coordinate.Y)
		time.Sleep(time.Duration(15) * time.Second)
	}
	// 钉钉工作台
	AdbShellInputTap(config.WorkPlatform.Coordinate.X, config.WorkPlatform.Coordinate.Y)
	time.Sleep(time.Duration(10) * time.Second)
	// 钉钉打卡位置并打卡
	AdbShellInputTap(config.Sign.Coordinate.X, config.Sign.Coordinate.Y)
	time.Sleep(time.Duration(15) * time.Second)
	AdbShellInputTap(config.Sign.Click.X, config.Sign.Click.Y)
	// 需要拍照
	if config.Sign.Photo.Require {
		AdbShellInputTap(config.Sign.Photo.Coordinate.X, config.Sign.Photo.Coordinate.Y)
		time.Sleep(time.Duration(5) * time.Second)
		AdbShellInputTap(config.Sign.Photo.Coordinate.X, config.Sign.Photo.Coordinate.Y)
	}
	time.Sleep(time.Duration(15) * time.Second)
	fmt.Println("钉钉打卡结束~")
	AdbShellInputKeyEvent("26")
	os.Exit(0)
}
