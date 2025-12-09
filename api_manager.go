package wukong_go_sdk

import (
	"context"
	"net/http"
)

// ManagerService 管理员相关接口
type ManagerService struct {
	client *Client
}

// ManagerLoginRequest 管理员登录请求
type ManagerLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ManagerUser 管理员用户信息
type ManagerUser struct {
	Username    string   `json:"username"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

// ManagerLoginResponse 管理员登录响应
type ManagerLoginResponse struct {
	Token  string      `json:"token"`
	Expire int64       `json:"expire"`
	User   ManagerUser `json:"user"`
}

// Login 管理员登录
// POST /manager/login
func (s *ManagerService) Login(ctx context.Context, req *ManagerLoginRequest) (*ManagerLoginResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody ManagerLoginResponse
	_, err := s.client.do(ctx, http.MethodPost, "/manager/login", req, &respBody)
	if err != nil {
		return nil, wrapError("manager.Login", err)
	}
	return &respBody, nil
}
