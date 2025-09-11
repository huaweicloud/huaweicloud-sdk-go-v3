package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetAuditAlarmLogStatusResponse Response Object
type BatchSetAuditAlarmLogStatusResponse struct {

	// 请求结果  - request_success: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchSetAuditAlarmLogStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAuditAlarmLogStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchSetAuditAlarmLogStatusResponse", string(data)}, " ")
}
