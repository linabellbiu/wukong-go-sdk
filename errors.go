package wukong_go_sdk

import pkgerrors "github.com/pkg/errors"

// APIError 表示 WuKongIM REST 接口的标准错误响应
// 参考文档示例：{"msg": "channel_id 参数不能为空", "status": 400}
type APIError struct {
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	if e.Msg != "" {
		return e.Msg
	}

	if e.Message != "" {
		return e.Message
	}

	return "wukongim api error"
}

// wrapError 统一包装底层错误，带上模块关键字和操作名
func wrapError(op string, err error) error {
	if err == nil {
		return nil
	}
	return pkgerrors.Wrap(err, "wukongimsdk err: "+op)
}
