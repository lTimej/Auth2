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


#### auth2简介
OAuth2提供了Access Token来解决授权第三方客户端访问受保护资源的问题；OIDC在这个基础上提供了ID Token来解决第三方客户端标识用户身份认证的问题。OIDC的核心在于在OAuth2的授权流程中，一并提供用户的身份认证信息（ID Token）给到第三方客户端，ID Token使用JWT格式来包装，得益于JWT（JSON Web Token）的自包含性，紧凑性以及防篡改机制，使得ID Token可以安全的传递给第三方客户端程序并且容易被验证。此外还提供了UserInfo的接口，用户获取用户的更完整的信息。


#### OIDC 主要术语
EU：End User：一个人类用户。
RP：Relying Party ,用来代指OAuth2中的受信任的客户端，身份认证和授权信息的消费方；
OP：OpenID Provider，有能力提供EU认证的服务（比如OAuth2中的授权服务），用来为RP提供EU的身份认证信息；
ID Token：JWT格式的数据，包含EU身份认证的信息。
UserInfo Endpoint：用户信息接口（受OAuth2保护），当RP使用Access Token访问时，返回授权用户的信息，此接口必须使用HTTPS。


#### OIDC 工作流程
RP发送一个认证请求给OP；
OP对EU进行身份认证，然后提供授权；
OP把ID Token和Access Token（需要的话）返回给RP；
RP使用Access Token发送一个请求UserInfo EndPoint；
UserInfo EndPoint返回EU的Claims。