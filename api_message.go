package wukong_go_sdk

import (
	"context"
	"net/http"
)

// MessageService 消息相关接口
type MessageService struct {
	client *Client
}

// MessageHeader 对应发送消息时的 header 字段
type MessageHeader struct {
	NoPersist int `json:"no_persist"`
	RedDot    int `json:"red_dot"`
	SyncOnce  int `json:"sync_once"`
}

// SendMessageRequest 发送消息请求体
type SendMessageRequest struct {
	Header      *MessageHeader `json:"header,omitempty"`
	ClientMsgNo string         `json:"client_msg_no"`
	FromUID     string         `json:"from_uid"`
	ChannelID   string         `json:"channel_id"`
	ChannelType ChannelType    `json:"channel_type"`
	Expire      int64          `json:"expire"`
	Payload     string         `json:"payload"`
	TagKey      string         `json:"tag_key,omitempty"`
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	MessageID   int64  `json:"message_id"`
	MessageSeq  int64  `json:"message_seq"`
	ClientMsgNo string `json:"client_msg_no"`
}

// SendMessage 发送单条消息
// POST /message/send
func (s *MessageService) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody SendMessageResponse
	_, err := s.client.do(ctx, http.MethodPost, "/message/send", req, &respBody)
	if err != nil {
		return nil, wrapError("message.SendMessage", err)
	}
	return &respBody, nil
}

// BatchSendMessageRequest 批量发送消息请求
type BatchSendMessageRequest struct {
	Messages []SendMessageRequest
}

// BatchSendMessageResponseItem 批量发送消息响应中的一项
type BatchSendMessageResponseItem struct {
	MessageID   int64  `json:"message_id"`
	MessageSeq  int64  `json:"message_seq"`
	ClientMsgNo string `json:"client_msg_no"`
}

// BatchSendMessage 批量发送消息
// POST /message/sendbatch
func (s *MessageService) BatchSendMessage(ctx context.Context, req *BatchSendMessageRequest) ([]BatchSendMessageResponseItem, error) {
	if req == nil {
		return nil, nil
	}

	var respBody []BatchSendMessageResponseItem
	_, err := s.client.do(ctx, http.MethodPost, "/message/sendbatch", req.Messages, &respBody)
	if err != nil {
		return nil, wrapError("message.BatchSendMessage", err)
	}
	return respBody, nil
}

// Message 通用消息结构
type Message struct {
	MessageID   int64       `json:"message_id"`
	MessageSeq  int64       `json:"message_seq"`
	ClientMsgNo string      `json:"client_msg_no"`
	FromUID     string      `json:"from_uid"`
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	Timestamp   int64       `json:"timestamp"`
	Payload     string      `json:"payload"`
}

// MessageSyncRequest 同步频道历史消息请求
type MessageSyncRequest struct {
	LoginUID        string      `json:"login_uid"`
	ChannelID       string      `json:"channel_id"`
	ChannelType     ChannelType `json:"channel_type"`
	StartMessageSeq int64       `json:"start_message_seq"`
	EndMessageSeq   int64       `json:"end_message_seq"`
	Limit           int         `json:"limit"`
	PullMode        int         `json:"pull_mode"`
}

// MessageSync 同步频道历史消息
// POST /channel/messagesync
func (s *MessageService) MessageSync(ctx context.Context, req *MessageSyncRequest) ([]Message, error) {
	if req == nil {
		return nil, nil
	}

	var respBody []Message
	_, err := s.client.do(ctx, http.MethodPost, "/channel/messagesync", req, &respBody)
	if err != nil {
		return nil, wrapError("message.MessageSync", err)
	}
	return respBody, nil
}

// MaxMessageSeqRequest 获取频道最大消息序号请求
type MaxMessageSeqRequest struct {
	ChannelID   string
	ChannelType ChannelType
}

// MaxMessageSeqResponse 获取频道最大消息序号响应
type MaxMessageSeqResponse struct {
	MaxMessageSeq int64 `json:"max_message_seq"`
}

// GetMaxMessageSeq 获取频道最大消息序号
// GET /channel/max_message_seq
func (s *MessageService) GetMaxMessageSeq(ctx context.Context, req *MaxMessageSeqRequest) (*MaxMessageSeqResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody MaxMessageSeqResponse
	_, err := s.client.do(ctx, http.MethodGet, "/channel/max_message_seq", nil, &respBody)
	if err != nil {
		return nil, wrapError("message.GetMaxMessageSeq", err)
	}
	return &respBody, nil
}

// UserSearchRequest 用户消息搜索请求
type UserSearchRequest struct {
	UID          string         `json:"uid"`
	Payload      map[string]any `json:"payload"`
	PayloadTypes []int          `json:"payload_types"`
	ChannelType  ChannelType    `json:"channel_type"`
	Limit        int            `json:"limit"`
	Page         int            `json:"page"`
	Highlights   []string       `json:"highlights"`
}

// UserSearchResponse 用户消息搜索响应
type UserSearchResponse struct {
	Total    int                `json:"total"`
	Limit    int                `json:"limit"`
	Page     int                `json:"page"`
	Messages []UserSearchResult `json:"messages"`
}

// UserSearchResult 用户消息搜索结果中的消息
type UserSearchResult struct {
	MessageID    int64          `json:"message_id"`
	MessageIDStr string         `json:"message_idstr"`
	MessageSeq   int64          `json:"message_seq"`
	ClientMsgNo  string         `json:"client_msg_no"`
	FromUID      string         `json:"from_uid"`
	ChannelID    string         `json:"channel_id"`
	ChannelType  ChannelType    `json:"channel_type"`
	Payload      map[string]any `json:"payload"`
	Topic        string         `json:"topic"`
	Timestamp    int64          `json:"timestamp"`
}

// UserSearch 用户消息搜索
// POST /plugins/wk.plugin.search/usersearch
func (s *MessageService) UserSearch(ctx context.Context, req *UserSearchRequest) (*UserSearchResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody UserSearchResponse
	_, err := s.client.do(ctx, http.MethodPost, "/plugins/wk.plugin.search/usersearch", req, &respBody)
	if err != nil {
		return nil, wrapError("message.UserSearch", err)
	}
	return &respBody, nil
}

// BatchSearchRequest 批量消息搜索请求
type BatchSearchRequest struct {
	MessageIDs []int64 `json:"message_ids"`
}

// BatchSearch 批量消息搜索
// POST /messages
func (s *MessageService) BatchSearch(ctx context.Context, req *BatchSearchRequest) ([]Message, error) {
	if req == nil {
		return nil, nil
	}

	var respBody []Message
	_, err := s.client.do(ctx, http.MethodPost, "/messages", req, &respBody)
	if err != nil {
		return nil, wrapError("message.BatchSearch", err)
	}
	return respBody, nil
}

// SingleSearchRequest 单条消息搜索请求
type SingleSearchRequest struct {
	MessageID int64 `json:"message_id"`
}

// SingleSearch 单条消息搜索
// POST /message
func (s *MessageService) SingleSearch(ctx context.Context, req *SingleSearchRequest) (*Message, error) {
	if req == nil {
		return nil, nil
	}

	var respBody Message
	_, err := s.client.do(ctx, http.MethodPost, "/message", req, &respBody)
	if err != nil {
		return nil, wrapError("message.SingleSearch", err)
	}
	return &respBody, nil
}
