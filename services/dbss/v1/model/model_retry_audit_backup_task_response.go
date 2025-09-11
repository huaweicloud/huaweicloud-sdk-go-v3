package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryAuditBackupTaskResponse Response Object
type RetryAuditBackupTaskResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RetryAuditBackupTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryAuditBackupTaskResponse struct{}"
	}

	return strings.Join([]string{"RetryAuditBackupTaskResponse", string(data)}, " ")
}
