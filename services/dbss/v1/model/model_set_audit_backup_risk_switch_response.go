package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditBackupRiskSwitchResponse Response Object
type SetAuditBackupRiskSwitchResponse struct {

	// 操作结果  - SUCCESS：成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditBackupRiskSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditBackupRiskSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetAuditBackupRiskSwitchResponse", string(data)}, " ")
}
