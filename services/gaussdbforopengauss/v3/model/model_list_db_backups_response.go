package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbBackupsResponse Response Object
type ListDbBackupsResponse struct {

	// 备份信息。
	Backups *[]BackupsResult `json:"backups,omitempty"`

	// 备份文件的总数。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListDbBackupsResponse", string(data)}, " ")
}
