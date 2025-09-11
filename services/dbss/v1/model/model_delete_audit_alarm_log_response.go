package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditAlarmLogResponse Response Object
type DeleteAuditAlarmLogResponse struct {

	// 请求结果  - request_success: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditAlarmLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditAlarmLogResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditAlarmLogResponse", string(data)}, " ")
}
