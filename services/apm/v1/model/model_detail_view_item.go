package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailViewItem 节点详情视图函数信息。
type DetailViewItem struct {

	// 表达式。
	Function *string `json:"function,omitempty"`

	// 作为。
	As *string `json:"as,omitempty"`
}

func (o DetailViewItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailViewItem struct{}"
	}

	return strings.Join([]string{"DetailViewItem", string(data)}, " ")
}
