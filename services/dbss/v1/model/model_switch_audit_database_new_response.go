package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAuditDatabaseNewResponse Response Object
type SwitchAuditDatabaseNewResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchAuditDatabaseNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAuditDatabaseNewResponse struct{}"
	}

	return strings.Join([]string{"SwitchAuditDatabaseNewResponse", string(data)}, " ")
}
