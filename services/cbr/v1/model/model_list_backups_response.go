package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsResponse Response Object
type ListBackupsResponse struct {

	// 备份列表
	Backups *[]BackupResp `json:"backups,omitempty"`

	// 备份个数
	Count *int32 `json:"count,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit          *int32 `json:"limit,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
