package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAllInstancesBackupsNewResponse struct {

	// 备份总数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	Backups        *[]Backup `json:"backups,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAllInstancesBackupsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllInstancesBackupsNewResponse struct{}"
	}

	return strings.Join([]string{"ShowAllInstancesBackupsNewResponse", string(data)}, " ")
}
