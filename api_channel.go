package wukong_go_sdk

import (
	"context"
	"fmt"
	"net/http"
)

// ChannelService 频道相关接口
type ChannelService struct {
	client *Client
}

// CreateChannelRequest 创建频道请求
type CreateChannelRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	Large       int         `json:"large"`
	Ban         int         `json:"ban"`
	Subscribers []string    `json:"subscribers"`
}

// CreateChannelResponse 创建频道响应
type CreateChannelResponse struct {
	Status string `json:"status"`
}

// Create 创建频道
// POST /channel
func (s *ChannelService) Create(ctx context.Context, req *CreateChannelRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.Create", err)
	}
	return &respBody, nil
}

// UpdateInfoRequest 更新频道信息请求
type UpdateInfoRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	Large       *int        `json:"large,omitempty"`
	Ban         *int        `json:"ban,omitempty"`
}

// UpdateInfo 更新频道信息
// POST /channel/info
func (s *ChannelService) UpdateInfo(ctx context.Context, req *UpdateInfoRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/info", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.UpdateInfo", err)
	}
	return &respBody, nil
}

// AddSubscribersRequest 添加频道订阅者请求
type AddSubscribersRequest struct {
	ChannelID      string      `json:"channel_id"`
	ChannelType    ChannelType `json:"channel_type"`
	Subscribers    []string    `json:"subscribers"`
	Reset          int         `json:"reset"`
	TempSubscriber int         `json:"temp_subscriber"`
}

// AddSubscribers 添加频道订阅者
// POST /channel/subscriber_add
func (s *ChannelService) AddSubscribers(ctx context.Context, req *AddSubscribersRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/subscriber_add", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.AddSubscribers", err)
	}
	return &respBody, nil
}

// RemoveSubscribersRequest 移除频道订阅者请求
type RemoveSubscribersRequest struct {
	ChannelID      string      `json:"channel_id"`
	ChannelType    ChannelType `json:"channel_type"`
	Subscribers    []string    `json:"subscribers"`
	TempSubscriber int         `json:"temp_subscriber"`
}

// RemoveSubscribers 移除频道订阅者
// POST /channel/subscriber_remove
func (s *ChannelService) RemoveSubscribers(ctx context.Context, req *RemoveSubscribersRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/subscriber_remove", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.RemoveSubscribers", err)
	}
	return &respBody, nil
}

// DeleteChannelRequest 删除频道请求
type DeleteChannelRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
}

// Delete 删除频道
// POST /channel/delete
func (s *ChannelService) Delete(ctx context.Context, req *DeleteChannelRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/delete", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.Delete", err)
	}
	return &respBody, nil
}

// ChannelUIDsRequest 通用 UID 列表请求
type ChannelUIDsRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	UIDs        []string    `json:"uids"`
}

// AddBlacklist 添加频道黑名单
// POST /channel/blacklist_add
func (s *ChannelService) AddBlacklist(ctx context.Context, req *ChannelUIDsRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/blacklist_add", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.AddBlacklist", err)
	}
	return &respBody, nil
}

// SetBlacklist 设置频道黑名单（替换）
// POST /channel/blacklist_set
func (s *ChannelService) SetBlacklist(ctx context.Context, req *ChannelUIDsRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/blacklist_set", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.SetBlacklist", err)
	}
	return &respBody, nil
}

// RemoveBlacklistRequest 移除频道黑名单请求
type RemoveBlacklistRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	UIDs        []string    `json:"uids"`
}

// RemoveBlacklist 移除频道黑名单
// POST /channel/blacklist_remove
func (s *ChannelService) RemoveBlacklist(ctx context.Context, req *RemoveBlacklistRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/blacklist_remove", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.RemoveBlacklist", err)
	}
	return &respBody, nil
}

// AddWhitelist 添加频道白名单
// POST /channel/whitelist_add
func (s *ChannelService) AddWhitelist(ctx context.Context, req *ChannelUIDsRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/whitelist_add", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.AddWhitelist", err)
	}
	return &respBody, nil
}

// SetWhitelist 设置频道白名单（替换）
// POST /channel/whitelist_set
func (s *ChannelService) SetWhitelist(ctx context.Context, req *ChannelUIDsRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/whitelist_set", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.SetWhitelist", err)
	}
	return &respBody, nil
}

// RemoveWhitelistRequest 移除频道白名单请求
type RemoveWhitelistRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	UIDs        []string    `json:"uids"`
}

// RemoveWhitelist 移除频道白名单
// POST /channel/whitelist_remove
func (s *ChannelService) RemoveWhitelist(ctx context.Context, req *RemoveWhitelistRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/whitelist_remove", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.RemoveWhitelist", err)
	}
	return &respBody, nil
}

// GetWhitelistRequest 获取频道白名单请求
type GetWhitelistRequest struct {
	ChannelID   string
	ChannelType ChannelType
}

// GetWhitelist 获取频道白名单
// GET /channel/whitelist
func (s *ChannelService) GetWhitelist(ctx context.Context, req *GetWhitelistRequest) ([]string, error) {
	if req == nil {
		return nil, nil
	}

	// 这里直接使用文档中的查询参数形式
	// /channel/whitelist?channel_id=xxx&channel_type=2
	path := "/channel/whitelist?channel_id=" + req.ChannelID + "&channel_type=" + fmt.Sprint(req.ChannelType)

	var respBody []string
	_, err := s.client.do(ctx, http.MethodGet, path, nil, &respBody)
	if err != nil {
		return nil, wrapError("channel.GetWhitelist", err)
	}
	return respBody, nil
}

// SetTmpSubscriberRequest 设置临时频道订阅者请求
type SetTmpSubscriberRequest struct {
	ChannelID   string      `json:"channel_id"`
	ChannelType ChannelType `json:"channel_type"`
	Subscribers []string    `json:"subscribers"`
}

// SetTmpSubscriber 设置临时频道订阅者
// POST /channel/tmp_subscriber_set
func (s *ChannelService) SetTmpSubscriber(ctx context.Context, req *SetTmpSubscriberRequest) (*CreateChannelResponse, error) {
	if req == nil {
		return nil, nil
	}

	var respBody CreateChannelResponse
	_, err := s.client.do(ctx, http.MethodPost, "/channel/tmp_subscriber_set", req, &respBody)
	if err != nil {
		return nil, wrapError("channel.SetTmpSubscriber", err)
	}
	return &respBody, nil
}
