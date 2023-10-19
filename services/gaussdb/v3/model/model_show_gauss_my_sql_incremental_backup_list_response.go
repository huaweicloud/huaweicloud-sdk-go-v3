package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlIncrementalBackupListResponse Response Object
type ShowGaussMySqlIncrementalBackupListResponse struct {

	// 备份信息。
	Backups *[]IncrementalBackups `json:"backups,omitempty"`

	// 备份文件的总数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGaussMySqlIncrementalBackupListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlIncrementalBackupListResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlIncrementalBackupListResponse", string(data)}, " ")
}
