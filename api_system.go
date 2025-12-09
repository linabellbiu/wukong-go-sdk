package wukong_go_sdk

import (
	"context"
	"net/http"
)

// SystemService 系统相关接口
type SystemService struct {
	client *Client
}

// HealthStatus 健康检查响应结构，根据文档示例只定义常见字段
// 如果文档有更详细字段，你可以在此基础上自行扩展。
type HealthStatus struct {
	Status string `json:"status"`
}

// Health 健康检查
// GET /health
func (s *SystemService) Health(ctx context.Context) (*HealthStatus, error) {
	var respBody HealthStatus
	_, err := s.client.do(ctx, http.MethodGet, "/health", nil, &respBody)
	if err != nil {
		return nil, wrapError("system.Health", err)
	}
	return &respBody, nil
}
