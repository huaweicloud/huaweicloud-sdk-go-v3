package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeAiAssistantReq AI助手订阅用户请求，支持添加或删除用户、用户组、项目的订阅。
type SubscribeAiAssistantReq struct {
	Add *SubscribeOperationReq `json:"add,omitempty"`

	Delete *SubscribeOperationReq `json:"delete,omitempty"`
}

func (o SubscribeAiAssistantReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAiAssistantReq struct{}"
	}

	return strings.Join([]string{"SubscribeAiAssistantReq", string(data)}, " ")
}
