package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllInstancesBackupsResponse Response Object
type ShowAllInstancesBackupsResponse struct {

	// 总记录数。
	TotalCount *int64 `json:"total_count,omitempty"`

	// 备份列表信息。
	Backups        *[]QueryInstanceBackupResponseBodyBackups `json:"backups,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowAllInstancesBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllInstancesBackupsResponse struct{}"
	}

	return strings.Join([]string{"ShowAllInstancesBackupsResponse", string(data)}, " ")
}
