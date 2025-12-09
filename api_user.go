package wukong_go_sdk

import (
	"context"
	"net/http"
)

// UserService 用户相关接口
type UserService struct {
	client *Client
}

// UpdateUserTokenRequest 更新用户 Token 请求
type UpdateUserTokenRequest struct {
	UID         string `json:"uid"`
	Token       string `json:"token"`
	DeviceFlag  int    `json:"device_flag"`
	DeviceLevel int    `json:"device_level"`
}

// UpdateToken 更新用户 Token
// POST /user/token
func (s *UserService) UpdateToken(ctx context.Context, req *UpdateUserTokenRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/user/token", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// DeviceQuitRequest 强制设备退出请求
type DeviceQuitRequest struct {
	UID        string `json:"uid"`
	DeviceFlag int    `json:"device_flag"`
}

// DeviceQuit 强制设备退出
// POST /user/device_quit
func (s *UserService) DeviceQuit(ctx context.Context, req *DeviceQuitRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/user/device_quit", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// OnlineStatusRequest 获取用户在线状态请求
type OnlineStatusRequest struct {
	UIDs []string
}

// UserOnlineStatus 单个用户在线状态
type UserOnlineStatus struct {
	UID        string       `json:"uid"`
	Online     OnlineStatus `json:"online"`
	DeviceFlag int          `json:"device_flag"`
}

// OnlineStatus 获取用户在线状态
// POST /user/onlinestatus
func (s *UserService) OnlineStatus(ctx context.Context, req *OnlineStatusRequest) ([]UserOnlineStatus, error) {
	if req == nil {
		return nil, nil
	}

	var respBody []UserOnlineStatus
	_, err := s.client.do(ctx, http.MethodPost, "/user/onlinestatus", req.UIDs, &respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// SystemUIDs 获取系统用户 ID 列表
// GET /user/systemuids
func (s *UserService) SystemUIDs(ctx context.Context) ([]string, error) {
	var respBody []string
	_, err := s.client.do(ctx, http.MethodGet, "/user/systemuids", nil, &respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// SystemUIDsChangeRequest 添加或移除系统用户 ID 请求
type SystemUIDsChangeRequest struct {
	UIDs []string `json:"uids"`
}

// AddSystemUIDs 添加系统用户 ID
// POST /user/systemuids_add
func (s *UserService) AddSystemUIDs(ctx context.Context, req *SystemUIDsChangeRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/user/systemuids_add", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}

// RemoveSystemUIDs 移除系统用户 ID
// POST /user/systemuids_remove
func (s *UserService) RemoveSystemUIDs(ctx context.Context, req *SystemUIDsChangeRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/user/systemuids_remove", req, &respBody)
	if err != nil {
		return nil, err
	}
	return &respBody, nil
}
