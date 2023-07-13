package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogsTreeList struct {

	// 文件日志树
	Trees *[]LogsTree `json:"trees,omitempty"`

	// 记录总数
	Total *int32 `json:"total,omitempty"`
}

func (o LogsTreeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogsTreeList struct{}"
	}

	return strings.Join([]string{"LogsTreeList", string(data)}, " ")
}
