package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LineViewItem 连接线视图函数详情。
type LineViewItem struct {

	// 表达式。
	Function *string `json:"function,omitempty"`

	// 作为。
	As *string `json:"as,omitempty"`
}

func (o LineViewItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineViewItem struct{}"
	}

	return strings.Join([]string{"LineViewItem", string(data)}, " ")
}
