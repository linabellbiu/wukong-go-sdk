package wukong_go_sdk

// ChannelType 表示频道类型，例如个人频道、群组频道等
type ChannelType int

const (
	// ChannelTypePerson 个人频道
	ChannelTypePerson ChannelType = 1
	// ChannelTypeGroup 群组频道
	ChannelTypeGroup ChannelType = 2
	// ChannelTypeCustomerService 客服频道
	ChannelTypeCustomerService ChannelType = 3
	// ChannelTypeCommunity 社区频道
	ChannelTypeCommunity ChannelType = 4
	// ChannelTypeCommunityTopic 社区话题频道
	ChannelTypeCommunityTopic ChannelType = 5
	// ChannelTypeNews 资讯频道
	ChannelTypeNews ChannelType = 6
	// ChannelTypeLivestream 直播频道
	ChannelTypeLivestream ChannelType = 9
	// ChannelTypeVisitor 访客频道
	ChannelTypeVisitor ChannelType = 10
	// ChannelTypeSingleAgent 单聊 Agent 频道
	ChannelTypeSingleAgent ChannelType = 11
	// ChannelTypeGroupAgent 群聊 Agent 频道
	ChannelTypeGroupAgent ChannelType = 12
)

// IntranetType 表示路由地址类型
// 0 表示外网地址，1 表示内网地址
type IntranetType int

const (
	// IntranetTypeExternal 外网地址
	IntranetTypeExternal IntranetType = 0
	// IntranetTypeInternal 内网地址
	IntranetTypeInternal IntranetType = 1
)

// OnlineStatus 表示用户在线状态
type OnlineStatus int

const (
	// OnlineStatusOffline 离线
	OnlineStatusOffline OnlineStatus = 0
	// OnlineStatusOnline 在线
	OnlineStatusOnline OnlineStatus = 1
)

// OnlyUnreadMode 表示会话同步时是否只返回有未读的会话
type OnlyUnreadMode int

const (
	// OnlyUnreadAll 返回所有会话
	OnlyUnreadAll OnlyUnreadMode = 0
	// OnlyUnreadUnread 仅返回有未读消息的会话
	OnlyUnreadUnread OnlyUnreadMode = 1
)
