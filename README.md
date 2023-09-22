# Auth2
auth2 client

# 项目结构
```
├── app             // 应用目录
│   ├── controller  // 控制器
│   ├── dao         // DAO层
│   ├── model       // 模型层
│   └── service     // 服务层
├── config          // 系统配置
├── docker
├── doc             // 文档中心
├── router          // 路由
├── utils           // 系统工具
├── go.mod
└── main.go
```

# swagger 文档
go get  github.com/swaggo/swag/cmd/swag@latest

cd /root/go/pkg/mod/github.com/swaggo/swag@v1.16.2/cmd/swag
go build
mv swag /usr/local/go/bin

### 执行完在项目目录生成docs文件夹
./docs
├── docs.go
├── swagger.json
└── swagger.yaml

swag init