package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditBackupSwitchResponse Response Object
type SetAuditBackupSwitchResponse struct {

	// 操作结果  - SUCCESS：成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditBackupSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditBackupSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetAuditBackupSwitchResponse", string(data)}, " ")
}
