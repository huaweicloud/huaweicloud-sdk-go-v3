package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAuditInstanceNewResponse Response Object
type StopAuditInstanceNewResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopAuditInstanceNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAuditInstanceNewResponse struct{}"
	}

	return strings.Join([]string{"StopAuditInstanceNewResponse", string(data)}, " ")
}
