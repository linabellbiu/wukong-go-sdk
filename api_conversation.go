package wukong_go_sdk

import (
	"context"
	"net/http"
)

// ConversationService 会话相关接口
type ConversationService struct {
	client *Client
}

// ConversationSyncRequest 同步用户会话请求
type ConversationSyncRequest struct {
	UID                 string         `json:"uid"`
	Version             int64          `json:"version"`
	LastMsgSeqs         string         `json:"last_msg_seqs"`
	MsgCount            int            `json:"msg_count"`
	OnlyUnread          OnlyUnreadMode `json:"only_unread"`
	ExcludeChannelTypes []ChannelType  `json:"exclude_channel_types"`
}

// ConversationRecentMessage 最近消息
type ConversationRecentMessage struct {
	MessageID   int64  `json:"message_id"`
	MessageSeq  int64  `json:"message_seq"`
	ClientMsgNo string `json:"client_msg_no"`
	FromUID     string `json:"from_uid"`
	Timestamp   int64  `json:"timestamp"`
	Payload     string `json:"payload"`
}

// Conversation 会话信息
type Conversation struct {
	ChannelID   string                      `json:"channel_id"`
	ChannelType ChannelType                 `json:"channel_type"`
	Unread      int                         `json:"unread"`
	Timestamp   int64                       `json:"timestamp"`
	LastMsgSeq  int64                       `json:"last_msg_seq"`
	Version     int64                       `json:"version"`
	Recents     []ConversationRecentMessage `json:"recents"`
}

// Sync 同步用户会话
// POST /conversation/sync
func (s *ConversationService) Sync(ctx context.Context, req *ConversationSyncRequest) ([]Conversation, error) {
	if req == nil {
		return nil, nil
	}

	var respBody []Conversation
	_, err := s.client.do(ctx, http.MethodPost, "/conversation/sync", req, &respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// ConversationClearUnreadRequest 清除未读消息请求
type ConversationClearUnreadRequest struct {
	UID         string      `json:"uid"`
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	MessageSeq  int64       `json:"message_seq"`
}

// ClearUnread 清除未读消息
// POST /conversations/clearUnread
func (s *ConversationService) ClearUnread(ctx context.Context, req *ConversationClearUnreadRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/conversations/clearUnread", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// ConversationSetUnreadRequest 设置会话未读数请求
type ConversationSetUnreadRequest struct {
	UID         string      `json:"uid"`
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	Unread      int         `json:"unread"`
}

// SetUnread 设置会话未读数
// POST /conversations/setUnread
func (s *ConversationService) SetUnread(ctx context.Context, req *ConversationSetUnreadRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/conversations/setUnread", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// ConversationDeleteRequest 删除会话请求
type ConversationDeleteRequest struct {
	UID         string      `json:"uid"`
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
}

// Delete 删除会话
// POST /conversations/delete
func (s *ConversationService) Delete(ctx context.Context, req *ConversationDeleteRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/conversations/delete", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
