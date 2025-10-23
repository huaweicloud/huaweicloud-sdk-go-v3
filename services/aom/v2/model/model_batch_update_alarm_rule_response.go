package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAlarmRuleResponse Response Object
type BatchUpdateAlarmRuleResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 资源列表。
	Resources      *[]BatchUpdateItemResult `json:"resources,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchUpdateAlarmRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAlarmRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateAlarmRuleResponse", string(data)}, " ")
}
