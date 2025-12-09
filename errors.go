package wukong_go_sdk

// APIError 表示 WuKongIM REST 接口的标准错误响应
// 参考文档示例：{"msg": "channel_id 参数不能为空", "status": 400}
type APIError struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Msg != "" {
		return e.Msg
	}
	return "wukongim api error"
}
