package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditBackupStatusResponse Response Object
type ShowAuditBackupStatusResponse struct {

	// 状态  - 1：成功 - 其他：失败
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAuditBackupStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditBackupStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditBackupStatusResponse", string(data)}, " ")
}
