package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInteractiveChatReq 创建交互助手对话请求。
type CreateInteractiveChatReq struct {

	// 角色ID。
	RoleId string `json:"role_id"`

	// 问题
	Message string `json:"message"`

	// 当前对话的唯一标识，用于关联对话上下文内容。
	SessionId string `json:"session_id"`

	Language *LanguageEnum `json:"language,omitempty"`
}

func (o CreateInteractiveChatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInteractiveChatReq struct{}"
	}

	return strings.Join([]string{"CreateInteractiveChatReq", string(data)}, " ")
}
