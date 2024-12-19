package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsDetailsResponse Response Object
type ListBackupsDetailsResponse struct {

	// 备份信息。
	Backups *[]ListBackupsResult `json:"backups,omitempty"`

	// 备份文件的总数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsDetailsResponse", string(data)}, " ")
}
