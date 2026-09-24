package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExpirationTimeResponse Response Object
type UpdateExpirationTimeResponse struct {

	// 成功修改过期时间的备份数量。
	AffectedBackupsCount *int32 `json:"affected_backups_count,omitempty"`

	// 修改后的备份过期时间，格式：YYYY-MM-DD。
	NewExpirationDay *string `json:"new_expiration_day,omitempty"`

	// 任务ID
	OperationLogId *string `json:"operation_log_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateExpirationTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpirationTimeResponse struct{}"
	}

	return strings.Join([]string{"UpdateExpirationTimeResponse", string(data)}, " ")
}
