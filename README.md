<div align="center"> <img width="444px" src="./assets/go_icon.png"/> </div>

# 支持的功能  
- 注册、登录、查看用户信息、更新用户信息、删除用户  

# 采用的开源项目：
<table >
    <tr>
      <th><img width="50px" src="./assets/gin.png"></th>
      <th><img width="50px" src="./assets/gorm.png"></th>
      <th><img width="50px" src="./assets/mysql.png"></th>
      <th><img width="50px" src="./assets/redis.png"></th>
    </tr>
    <tr>
      <td align="center"><a href="https://github.com/gin-gonic/gin">Gin</a></td>
      <td align="center"><a href="https://github.com/go-gorm/gorm">GORM</a></td>
      <td align="center"><a href="https://www.mysql.com/">MySql</a></td>
      <td align="center"><a href="https://github.com/redis/go-redis">Redis</a></td>
    </tr>
  </table>


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
