学习golang写的im即时通信项目，有单聊、群聊、图片等功能，基于go、vue、websocket。

# 安装方法

## 1.安装依赖
```bash
go get github.com/go-xorm/xorm
go get github.com/gorilla/websocket
go get gopkg.in/fatih/set.v0
go get github.com/go-xorm/xorm
```


## 2.项目配置

修改service/init.go 中数据库配置文件
```cgo
drivename :="mysql"
DsName := "root:root@(192.168.0.102:3306)/chat?charset=utf8"
```
为你自己的数据库以及密码,格式如下
```
用户名:密码@(ip:port)/数据库名称?charset=utf8
```
### 页面入口地址
```
http://127.0.0.1:8080/
```

### 3.示例截图
![文字](https://github.com/iScript/go-im-example/blob/master/5135211C-C3EB-45D9-A7E9-B2E9CCAF7C80.png)
![文字](https://github.com/iScript/go-im-example/blob/master/A0BAF1E0-BEDD-4DB8-AC68-B8E647062CBC.png)
