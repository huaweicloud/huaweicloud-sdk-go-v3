package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMessageClearRuleReq 消息清理规则请求体
type SetMessageClearRuleReq struct {

	// 最多保留记录数
	MessageRetainNumber *int32 `json:"message_retain_number,omitempty"`
}

func (o SetMessageClearRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMessageClearRuleReq struct{}"
	}

	return strings.Join([]string{"SetMessageClearRuleReq", string(data)}, " ")
}
