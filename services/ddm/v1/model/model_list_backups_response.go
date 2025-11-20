package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsResponse Response Object
type ListBackupsResponse struct {

	// 备份总数。
	Total *int32 `json:"total,omitempty"`

	// 分页参数起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 备份信息。
	Backups        *[]BackupInfo `json:"backups,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
