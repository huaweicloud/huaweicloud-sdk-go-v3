package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostOldRequestsReq struct {
	Extends *OldExtends `json:"extends,omitempty"`

	// 默认false true：使用内部闲聊语料进行兜底 false：不使用闲聊兜底
	ChatEnable *bool `json:"chat_enable,omitempty"`

	// true：启动内部阈值 返回经过阈值处理之后的答案。 false：不启用内部阈值 返回原答案。
	ThresholdEnable *bool `json:"threshold_enable,omitempty"`

	// 用户id，在日志中用于标识不通用户，可以为任意String
	UserId *string `json:"user_id,omitempty"`

	// 用户输入
	Question string `json:"question"`

	// 调用接口时候传入，用以标记的问答的行为，默认为0，最终会展示在问答日志里。 0 用户输入 1 单击热点问题 3 单击推荐问题 4 单击问题提示
	OperateType *int32 `json:"operate_type,omitempty"`

	// 会话标识符，UUID格式。如：c04e6f7b-61d7-4a2d-a0c8-f9ecd2f62359。  每次对话开启，机器人创建会话id，下次请求中传入该id表示继续该轮对话，每轮会话有效时间为2分钟。 若传入的会话id已过期或者为空，则机器人会重新创建新的会话id（重新创建会话id会消耗一定时间）。
	SessionId *string `json:"session_id,omitempty"`
}

func (o PostOldRequestsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostOldRequestsReq struct{}"
	}

	return strings.Join([]string{"PostOldRequestsReq", string(data)}, " ")
}
