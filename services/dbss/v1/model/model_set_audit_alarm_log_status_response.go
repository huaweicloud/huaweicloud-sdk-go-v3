package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditAlarmLogStatusResponse Response Object
type SetAuditAlarmLogStatusResponse struct {

	// 请求结果  - request_success：成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditAlarmLogStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditAlarmLogStatusResponse struct{}"
	}

	return strings.Join([]string{"SetAuditAlarmLogStatusResponse", string(data)}, " ")
}
