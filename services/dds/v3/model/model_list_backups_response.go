package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackupsResponse struct {

	// 备份列表。
	Backups *[]BackupForList `json:"backups,omitempty" xml:"backups"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
