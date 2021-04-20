<h1 align="center"> DingAutoSign </h1>

<p align="center">通过 adb 来实现自动打卡</p>

<p align="center"><b>仅用于学习和测试，作者本人并不对其负责，请于运行测试完成后自行删除，请勿滥用！</b></p>


### 使用说明
1. Windows操作系统和 go 环境
2. 需要安卓 adb 驱动，并且启动 ```adb start-server```
3. 一台安卓手机连接电脑，开启 usb 调试功能。查看是否有连接上命令 ``` adb devices ```
4. 将 temp_config.json 改名为 config.json
5. 需要定时执行的话, 可以用 windows 定时任务来执行 execute.bat

### 配置须知
```text
{
  // 锁屏配置 
  "lock": {
    "require": true, // 是否有锁屏, true为有
    "password": "****", // 锁屏密码
    "coordinate": {
      "x": 0, // 锁屏解锁坐标X
      "y": 0  // 锁屏解锁坐标X
    }
  },
  // 解锁后是否需要启动钉钉 
  "start_dingding": {
    "require": true, // 需要则为true 
    "coordinate": {
      "x": 0, // 钉钉应用在桌面坐标X
      "y": 0 // 钉钉应用在桌面坐标Y
    }
  },
  // 钉钉工作台
  "work_platform": {
    "coordinate": {
      "x": 553, // 钉钉工作台坐标X
      "y": 1703 // 钉钉工作台坐标Y
    }
  },
  // 钉钉打卡坐标
  "sign": {
    "coordinate": {
      "x": 0, // 钉钉打卡坐标X
      "y": 0  // 钉钉打卡坐标Y
    },
    // 点击钉钉打卡
    "click": {
      "x": 0, // 点击钉钉打卡坐标X
      "y": 0  // 点击钉钉打卡坐标Y
    },
    // 钉钉打卡需要拍照
    "photo": {
      "require": false, // 需要则为true
      "coordinate": {
        "x": 0, // 钉钉打卡拍照坐标X
        "y": 0  // 钉钉打卡拍照坐标Y
      }
    }
  }
}
```
## 参考
* [adb常用命令(golang版）及输入中文](https://blog.csdn.net/weixin_30635053/article/details/96171154)
* [ADB 操作命令详解及用法大全](https://juejin.cn/post/6844903645289398280#heading-32)

## License
MIT