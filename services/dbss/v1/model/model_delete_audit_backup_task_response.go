package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditBackupTaskResponse Response Object
type DeleteAuditBackupTaskResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditBackupTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditBackupTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditBackupTaskResponse", string(data)}, " ")
}
