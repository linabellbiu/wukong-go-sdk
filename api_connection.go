package wukong_go_sdk

import (
	"context"
	"net/http"
)

// ConnectionService 连接相关接口
type ConnectionService struct {
	client *Client
}

// ConnectionRequest 移除/踢出连接请求
type ConnectionRequest struct {
	UID    string `json:"uid"`
	ConnID int64  `json:"conn_id"`
	NodeID int64  `json:"node_id"`
}

// Remove 移除连接
// POST /conn/remove
func (s *ConnectionService) Remove(ctx context.Context, req *ConnectionRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/conn/remove", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// Kick 踢出连接
// POST /conn/kick
func (s *ConnectionService) Kick(ctx context.Context, req *ConnectionRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/conn/kick", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
