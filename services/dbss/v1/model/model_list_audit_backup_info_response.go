package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditBackupInfoResponse Response Object
type ListAuditBackupInfoResponse struct {

	// 备份列表
	DataList *[]BackupInfo `json:"data_list,omitempty"`

	// 总记录数
	TotalNum       *int64 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditBackupInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditBackupInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAuditBackupInfoResponse", string(data)}, " ")
}
