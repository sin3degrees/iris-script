# iris-script

```
conf  配置文件
controllers  控制器 入参处理 api的入口
datasource 数据库配置 
models  结构体
db  sql数据文件 postman接口文件
repo 数据库的操作
middleware 中间件 jwt实现
route  注册路由
service 业务逻辑代码
utils  工具类
config.json 配置文件的映射
main.go 主程序入口
```
### 启动项目
```
1.安装依赖 go get
2.go run main.go
```
1. 使用go get直接下载依赖，或在github手动下载包放到gopath/src/github.com/
2. 导包时使用相对路径需要将项目放在你配置的GOPATH目录下
3. 使用go mod init用go mod来管理 
