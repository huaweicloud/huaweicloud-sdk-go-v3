package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubfilesResponse struct {

	// 文件日志树
	Trees *[]LogsTree `json:"trees,omitempty"`

	// 记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubfilesResponse struct{}"
	}

	return strings.Join([]string{"ListSubfilesResponse", string(data)}, " ")
}
