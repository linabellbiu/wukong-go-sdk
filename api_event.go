package wukong_go_sdk

import (
	"context"
	"fmt"
	"net/http"
)

// EventService 事件相关接口
type EventService struct {
	client *Client
}

// EventPayload 事件负载
type EventPayload struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

// EventSendRequest 发送事件请求
type EventSendRequest struct {
	ClientMsgNo string       `json:"client_msg_no"`
	ChannelID   string       `json:"channel_id"`
	ChannelType ChannelType  `json:"channel_type"`
	FromUID     string       `json:"from_uid"`
	Event       EventPayload `json:"event"`

	// ForceEnd 是否强制结束现有流，对应 force_end 查询参数，可选
	ForceEnd *int `json:"-"`
}

// Send 发送事件
// POST /event
func (s *EventService) Send(ctx context.Context, req *EventSendRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	path := "/event"
	if req.ForceEnd != nil {
		path = fmt.Sprintf("/event?force_end=%d", *req.ForceEnd)
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, path, req, &respBody)
	if err != nil {
		return nil, wrapError("event.Send", err)
	}
	return &respBody, nil
}
