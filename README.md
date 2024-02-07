# 关于 
gin-demo基于gin及开源项目进行简单且完善的API开发。  
本项目致力于帮助初学者掌握Go WEB API开发!  
本项目保持简单！  
本项目仅供学习参考！  

主要使用以下第三方包：
- [gin](https://github.com/gin-gonic/gin) 项目路由
- [gorm](https://github.com/go-gorm/gorm) Go语言的Mysql关系型ORM
- [go-redis](https://github.com/redis/go-redis) Go语言的redis客户端
- [viper](https://github.com/spf13/viper) 配置文件解析
- [cors](https://github.com/gin-contrib/cors) 网站请求API，跨域处理
- [jwt-go](https://github.com/golang-jwt/jwt) 用户的身份校验
- [zap](https://github.com/uber-go/zap) 日志记录
- [lumberjack](https://github.com/natefinch/lumberjack) 日志切割

# 项目结构
项目启动：`go mod tidy && go run main.go`
```
├── README.md
├── api_test.go  //测试文件
├── docs
│   └── db.sql // 数据库文件
├── etc
│   └── config.yaml //配置文件
├── go.mod // 第三方依赖库
├── go.sum
├── internal // 项目内部文件夹
│   ├── api
│   │   └── user.go // API定义文件
│   ├── api.go // 项目启动引擎文件
│   ├── cache
│   │   └── user.go // redis缓存文件
│   ├── config
│   │   └── config.go // 配置文件解析结构文件
│   ├── logic // 业务逻辑处理文件夹
│   │   └── user.go 
│   ├── model // 数据库实体
│   │   └── user.go
│   └── router // 路由文件夹
│       ├── middleware // 路由中间件
│       └── route.go // 路由定义文件
├── main.go // 项目启动文件
└── pkg
    └── response.go // 自定义公用工具
```

# 文档&&学习交流
更多参考：[语雀gin-demo](https://www.yuque.com/ngyhd/sdqiox/iyosrxglvvbm5b36)
