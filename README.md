# Acosmi Desktop SDK for Go

**acosmi/desktop-sdk-go** 是创宇太虚 (Acosmi) 桌面智能体的 Go 语言 OAuth 授权 SDK。

它封装了完整的 OAuth 2.1 + PKCE 授权流程，让第三方桌面智能体可以安全地接入 Acosmi 的托管模型、权益系统等 API。

## 安装

```bash
go get github.com/acosmi/desktop-sdk-go
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    acosmi "github.com/acosmi/desktop-sdk-go"
)

func main() {
    // 创建客户端
    client, err := acosmi.NewClient(acosmi.Config{
        ServerURL: "http://127.0.0.1:3300/api/v4",
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
    defer cancel()

    // 首次使用需要登录 (会自动打开浏览器)
    if !client.IsAuthorized() {
        err := client.Login(ctx, "My Desktop Agent", []string{
            "models:chat",
            "entitlements",
        })
        if err != nil {
            log.Fatalf("授权失败: %v", err)
        }
    }

    // 获取可用模型
    models, err := client.ListModels(ctx)
    if err != nil {
        log.Fatalf("获取模型失败: %v", err)
    }
    for _, m := range models {
        fmt.Printf("模型: %s (%s)\n", m.Name, m.Provider)
    }
}
```

## 功能特性

| 功能 | 描述 |
|------|------|
| **OAuth 2.1 + PKCE** | 完整的授权码流程，无需 client_secret |
| **动态客户端注册** | RFC 7591，自动注册桌面客户端 |
| **Token 自动刷新** | access_token 过期前自动刷新，调用方无感 |
| **Token 持久化** | 内置文件存储，支持自定义实现（如系统钥匙串） |
| **托管模型调用** | 同步聊天 + SSE 流式聊天 |
| **权益查询** | 查询 Token 余额和调用次数 |
| **线程安全** | 所有 API 调用线程安全 |

## API 参考

### 客户端

```go
// 创建客户端
client, err := acosmi.NewClient(acosmi.Config{
    ServerURL:  "http://127.0.0.1:3300/api/v4",
    Store:      nil,  // 默认文件存储 ~/.acosmi/desktop-tokens.json
    HTTPClient: nil,  // 默认 30s 超时
})

// 授权
client.Login(ctx, "appName", []string{"models:chat"})

// 查询余额
balance, _ := client.GetBalance(ctx)

// 获取模型列表
models, _ := client.ListModels(ctx)

// 同步聊天
resp, _ := client.Chat(ctx, modelID, acosmi.ChatRequest{...})

// 流式聊天
eventCh, errCh := client.ChatStream(ctx, modelID, acosmi.ChatRequest{...})

// 登出
client.Logout(ctx)
```

### 自定义 Token 存储

实现 `TokenStore` 接口即可替换默认的文件存储：

```go
type TokenStore interface {
    Save(tokens *TokenSet) error
    Load() (*TokenSet, error)
    Clear() error
}
```

## 完整示例

请参考 [example/main.go](./example/main.go) 获取完整的使用示例。

## 支持平台

- macOS
- Linux
- Windows

## 许可证

PolyForm Noncommercial License 1.0.0
