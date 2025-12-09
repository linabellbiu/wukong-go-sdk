package wukong_go_sdk

import (
	"context"
	"fmt"
	"resty.dev/v3"
	"time"
)

// Config 是 SDK 的基础配置
// BaseURL 例如：http://localhost:5001
// Token / AppKey 等认证信息后续可按实际文档扩展
// Timeout 为每次请求的超时时间
// Debug 控制 resty 的调试输出
type Config struct {
	BaseURL string

	// 认证相关字段：根据实际 WuKongIM 文档再补充
	Token  string
	AppKey string

	Timeout time.Duration
	Debug   bool
}

// Client 是 WuKongIM API 的客户端
// 通过 NewClient 创建
type Client struct {
	cfg *Config
	cli *resty.Client

	Route        *RouteService
	Message      *MessageService
	Channel      *ChannelService
	User         *UserService
	Conversation *ConversationService
	Connection   *ConnectionService
	Event        *EventService
	Manager      *ManagerService
	System       *SystemService
}

// NewClient 根据配置创建一个新的 WuKongIM 客户端
func NewClient(cfg Config) *Client {
	c := resty.New()

	if cfg.BaseURL != "" {
		c.SetBaseURL(cfg.BaseURL)
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 10 * time.Second
	}
	c.SetTimeout(cfg.Timeout)

	c.SetHeader("Content-Type", "application/json")

	if cfg.Debug {
		c.SetDebug(true)
	}

	// 配置全局错误反序列化结构
	c.SetError(&APIError{})

	client := &Client{
		cfg: &cfg,
		cli: c,
	}

	// 统一设置认证（例如 Header 或 Query），具体根据 WuKongIM 文档调整
	client.applyAuth()

	// 初始化分组服务
	client.Route = &RouteService{client: client}
	client.Message = &MessageService{client: client}
	client.Channel = &ChannelService{client: client}
	client.User = &UserService{client: client}
	client.Conversation = &ConversationService{client: client}
	client.Connection = &ConnectionService{client: client}
	client.Event = &EventService{client: client}
	client.Manager = &ManagerService{client: client}
	client.System = &SystemService{client: client}

	return client
}

// applyAuth 根据配置把认证信息挂到 resty 客户端
// 这里先留一个通用实现，后续再根据文档完善
func (c *Client) applyAuth() {
	if c.cfg.Token != "" {
		// 使用 Resty 内置的认证配置，在每个请求上附加 Authorization 头
		c.cli.SetAuthScheme("Bearer")
		c.cli.SetAuthToken(c.cfg.Token)
	}
}

// do 执行 HTTP 请求的公共封装
// method: GET/POST/PUT/DELETE
// path: 相对路径，例如 "/api/v1/message/send"
// reqBody: 请求体结构体，会被编码为 JSON
// respBody: 响应体结构体指针，用于 JSON 反序列化
func (c *Client) do(ctx context.Context, method, path string, reqBody any, respBody any) (*resty.Response, error) {
	req := c.cli.R().SetContext(ctx)

	if reqBody != nil {
		req.SetBody(reqBody)
	}

	if respBody != nil {
		req.SetResult(respBody)
	}

	resp, err := req.Execute(method, path)
	if err != nil {
		return nil, wrapError("client.do", err)
	}

	if resp.IsError() {
		if errObj, ok := resp.Error().(*APIError); ok && (errObj.Msg != "" || errObj.Message != "" || errObj.Status != 0) {
			return resp, errObj
		}
		return resp, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return resp, nil
}
