package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostQaSessionReq struct {
	Extends *SessionExtends `json:"extends,omitempty"`
	// 默认true true：使用内部闲聊语料进行兜底 false：不使用闲聊兜底

	ChatEnable *bool `json:"chat_enable,omitempty"`
	// 用户id，在日志中用于标识不通用户，可以为任意String。

	UserId *string `json:"user_id,omitempty"`
	// 用户输入。

	Question string `json:"question"`
}

func (o PostQaSessionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostQaSessionReq struct{}"
	}

	return strings.Join([]string{"PostQaSessionReq", string(data)}, " ")
}
