package wukong_go_sdk

import (
	"context"
	"fmt"
	"net/http"
)

// RouteService 路由相关接口
type RouteService struct {
	client *Client
}

// RouteAddressRequest 请求 IM 地址时的查询参数
type RouteAddressRequest struct {
	// 0 - 返回外网地址；1 - 返回内网地址
	Intranet IntranetType
}

// RouteAddress 单个地址响应
type RouteAddress struct {
	TCPAddr string `json:"tcp_addr"`
	WSAddr  string `json:"ws_addr"`
	WSSAddr string `json:"wss_addr"`
}

// GetIMAddress 获取用户 IM 地址
// GET /route?intranet=0
func (s *RouteService) GetIMAddress(ctx context.Context, req *RouteAddressRequest) (*RouteAddress, error) {
	if req == nil {
		req = &RouteAddressRequest{}
	}

	path := "/route"
	if req.Intranet == IntranetTypeExternal || req.Intranet == IntranetTypeInternal {
		path = fmt.Sprintf("/route?intranet=%d", req.Intranet)
	}

	var respBody RouteAddress
	_, err := s.client.do(ctx, http.MethodGet, path, nil, &respBody)
	if err != nil {
		return nil, wrapError("route.GetIMAddress", err)
	}
	return &respBody, nil
}

// BatchRouteAddressRequest 批量获取 IM 地址请求
type BatchRouteAddressRequest struct {
	// 0 - 返回外网地址；1 - 返回内网地址
	Intranet IntranetType
	// 用户 UID 列表
	UIDs []string
}

// BatchRouteAddress 批量地址响应中的一项
type BatchRouteAddress struct {
	UIDs    []string `json:"uids"`
	TCPAddr string   `json:"tcp_addr"`
	WSAddr  string   `json:"ws_addr"`
	WSSAddr string   `json:"wss_addr"`
}

// BatchGetIMAddress 批量获取用户 IM 地址
// POST /route/batch?intranet=0
func (s *RouteService) BatchGetIMAddress(ctx context.Context, req *BatchRouteAddressRequest) ([]BatchRouteAddress, error) {
	if req == nil {
		return nil, nil
	}

	path := "/route/batch"
	if req.Intranet == IntranetTypeExternal || req.Intranet == IntranetTypeInternal {
		path = fmt.Sprintf("/route/batch?intranet=%d", req.Intranet)
	}

	var respBody []BatchRouteAddress
	_, err := s.client.do(ctx, http.MethodPost, path, req.UIDs, &respBody)
	if err != nil {
		return nil, wrapError("route.BatchGetIMAddress", err)
	}
	return respBody, nil
}
