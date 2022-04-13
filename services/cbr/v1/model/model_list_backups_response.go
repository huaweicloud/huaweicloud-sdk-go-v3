package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackupsResponse struct {
	// 备份列表

	Backups *[]BackupResp `json:"backups,omitempty"`
	// 备份个数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
