package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackupsResponse struct {

	// 备份列表
	Backups *[]BackupResp `json:"backups,omitempty" xml:"backups"`

	// 备份个数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit          *int32 `json:"limit,omitempty" xml:"limit"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
