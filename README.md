<h1 align="center"> DingAutoSign </h1>

<p align="center">通过adb来实现自动打卡, 仅限于代码学习与交流</p>

### 使用说明
1. 需要 go 环境
2. 需要安卓 adb 驱动
3. 将temp_config.json 改名为 config.json
4. 需要定时执行的话, 可以用windows定时任务来执行execute.bat

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