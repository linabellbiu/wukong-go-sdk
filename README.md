# wukong-go-sdk

基于 `resty.dev/v3` 的 WuKongIM HTTP API Go SDK。

当前 SDK 按功能分组封装了 WuKongIM 的常用 HTTP 接口：

- Route：路由 / IM 接入地址
- Message：消息发送、历史同步、搜索
- Channel：频道管理、成员、黑白名单、临时订阅者
- User：用户 token、设备退出、在线状态、系统用户 ID
- Conversation：会话同步与未读管理
- Connection：连接移除、踢出
- Event：事件发送
- Manager：管理员登录
- System：系统健康检查

---

## 安装

```bash
go get github.com/intelli-train-ai/wukong-go-sdk
```

Go 版本：建议 Go 1.23 及以上。

---

## 初始化 Client

```go
import (
	"context"
	"log"

	wukong "github.com/intelli-train-ai/wukong-go-sdk"
)

func main() {
	cli := wukong.NewClient(wukong.Config{
		BaseURL: "http://localhost:5001", // WuKongIM HTTP API 地址
		Token:   "your-token",            // 如服务端开启鉴权则必填
	})

	ctx := context.Background()

	addr, err := cli.Route.GetIMAddress(ctx, &wukong.RouteAddressRequest{
		Intranet: wukong.IntranetTypeExternal,
	})
	if err != nil {
		log.Fatalf("GetIMAddress err: %+v", err)
	}

	log.Println("tcp=", addr.TCPAddr, "ws=", addr.WSAddr, "wss=", addr.WSSAddr)
}
```

---

## API 分组与方法一览

### RouteService（路由）

- `GetIMAddress(ctx, req)`  
  - **GET** `/route?intranet=0|1`

- `BatchGetIMAddress(ctx, req)`  
  - **POST** `/route/batch?intranet=0|1`

---

### MessageService（消息）

- `SendMessage(ctx, req)`  
  - **POST** `/message/send`

- `BatchSendMessage(ctx, req)`  
  - **POST** `/message/sendbatch`

- `MessageSync(ctx, req)`  
  - **POST** `/channel/messagesync`

- `GetMaxMessageSeq(ctx, req)`  
  - **GET** `/channel/max_message_seq`

- `UserSearch(ctx, req)`  
  - **POST** `/plugins/wk.plugin.search/usersearch`

- `BatchSearch(ctx, req)`  
  - **POST** `/messages`

- `SingleSearch(ctx, req)`  
  - **POST** `/message`

---

### ChannelService（频道）

- `Create(ctx, req)`  
  - **POST** `/channel`

- `UpdateInfo(ctx, req)`  
  - **POST** `/channel/info`

- `AddSubscribers(ctx, req)`  
  - **POST** `/channel/subscriber_add`

- `RemoveSubscribers(ctx, req)`  
  - **POST** `/channel/subscriber_remove`

- `Delete(ctx, req)`  
  - **POST** `/channel/delete`

- `AddBlacklist(ctx, req)`  
  - **POST** `/channel/blacklist_add`

- `SetBlacklist(ctx, req)`  
  - **POST** `/channel/blacklist_set`

- `RemoveBlacklist(ctx, req)`  
  - **POST** `/channel/blacklist_remove`

- `AddWhitelist(ctx, req)`  
  - **POST** `/channel/whitelist_add`

- `SetWhitelist(ctx, req)`  
  - **POST** `/channel/whitelist_set`

- `RemoveWhitelist(ctx, req)`  
  - **POST** `/channel/whitelist_remove`

- `GetWhitelist(ctx, req)`  
  - **GET** `/channel/whitelist?channel_id=...&channel_type=...`

- `SetTmpSubscriber(ctx, req)`  
  - **POST** `/channel/tmp_subscriber_set`

---

### UserService（用户）

- `UpdateToken(ctx, req)`  
  - **POST** `/user/token`

- `DeviceQuit(ctx, req)`  
  - **POST** `/user/device_quit`

- `OnlineStatus(ctx, req)`  
  - **POST** `/user/onlinestatus`

- `SystemUIDs(ctx)`  
  - **GET** `/user/systemuids`

- `AddSystemUIDs(ctx, req)`  
  - **POST** `/user/systemuids_add`

- `RemoveSystemUIDs(ctx, req)`  
  - **POST** `/user/systemuids_remove`

---

### ConversationService（会话）

- `Sync(ctx, req)`  
  - **POST** `/conversation/sync`

- `ClearUnread(ctx, req)`  
  - **POST** `/conversations/clearUnread`

- `SetUnread(ctx, req)`  
  - **POST** `/conversations/setUnread`

- `Delete(ctx, req)`  
  - **POST** `/conversations/delete`

---

### ConnectionService（连接）

- `Remove(ctx, req)`  
  - **POST** `/conn/remove`

- `Kick(ctx, req)`  
  - **POST** `/conn/kick`

---

### EventService（事件）

- `Send(ctx, req)`  
  - **POST** `/event`（可选 `force_end` 查询参数）

---

### ManagerService（管理员）

- `Login(ctx, req)`  
  - **POST** `/manager/login`

---

### SystemService（系统）

- `Health(ctx)`  
  - **GET** `/health`

---

## 简单的错误处理约定

- SDK 内部使用 `github.com/pkg/errors` 对底层错误做包装，统一前缀为：

  `wukongimsdk err: <模块.方法>: <底层错误>`

- 当 WuKongIM 返回业务错误（JSON 中包含 `msg`、`status` 字段）时，会返回 `*APIError`：

  ```go
  type APIError struct {
  	Msg    string `json:"msg"`
  	Status int    `json:"status"`
  }
  ```

- 建议调用方这样区分：

  ```go
  resp, err := cli.Message.SendMessage(ctx, req)
  if err != nil {
  	if apiErr, ok := err.(*wukong.APIError); ok {
  		// WuKongIM 业务错误
  		log.Printf("wukongim api error: status=%d msg=%s", apiErr.Status, apiErr.Msg)
  		return
  	}

  	// SDK / 网络层错误
  	log.Printf("wukongimsdk err: %+v", err)
  	return
  }
  ```