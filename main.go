package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
