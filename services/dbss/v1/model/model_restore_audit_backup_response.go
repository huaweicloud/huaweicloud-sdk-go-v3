package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreAuditBackupResponse Response Object
type RestoreAuditBackupResponse struct {

	// 操作结果  - SUCCESS：成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreAuditBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreAuditBackupResponse struct{}"
	}

	return strings.Join([]string{"RestoreAuditBackupResponse", string(data)}, " ")
}
