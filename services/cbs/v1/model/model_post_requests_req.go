package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostRequestsReq struct {
	Extends *Extends `json:"extends,omitempty"`

	// 默认false true：使用内部闲聊语料进行兜底 false：不使用闲聊兜底
	ChatEnable *bool `json:"chat_enable,omitempty"`

	// 用户id，在日志中用于标识不通用户，可以为任意String
	UserId *string `json:"user_id,omitempty"`

	// 用户输入
	Question string `json:"question"`

	// 会话标识符，UUID格式。如：c04e6f7b-61d7-4a2d-a0c8-f9ecd2f62359。  每次对话开启，机器人创建会话id，下次请求中传入该id表示继续该轮对话，每轮会话有效时间为2分钟。 若传入的会话id已过期或者为空，则机器人会重新创建新的会话id（重新创建会话id会消耗一定时间）。
	SessionId *string `json:"session_id,omitempty"`

	// 指定发送的机器人类型集合。  0 知识库问答。  1 技能问答。  2 闲聊问答。  3 图谱问答。  4 文档问答。  5 表格问答。  非必填字段。如果不填，会使用默认的机器人融合策略。
	QueryTypes *[]int32 `json:"query_types,omitempty"`
}

func (o PostRequestsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostRequestsReq struct{}"
	}

	return strings.Join([]string{"PostRequestsReq", string(data)}, " ")
}
