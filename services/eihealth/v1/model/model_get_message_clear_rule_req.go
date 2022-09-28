package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息清理规则DTO
type GetMessageClearRuleReq struct {

	// 保存时长，单位：天
	MessageRetainTime *int32 `json:"message_retain_time,omitempty"`

	// 最多保留记录数
	MessageRetainNumber *int32 `json:"message_retain_number,omitempty"`
}

func (o GetMessageClearRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetMessageClearRuleReq struct{}"
	}

	return strings.Join([]string{"GetMessageClearRuleReq", string(data)}, " ")
}
