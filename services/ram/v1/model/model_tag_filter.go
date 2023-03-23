package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 根据标签筛选。
type TagFilter struct {

	// 标签\"键\"的标识符或名称。
	Key string `json:"key"`

	// 标签\"键\"对应的\"值\"列表。
	Values *[]string `json:"values,omitempty"`
}

func (o TagFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFilter struct{}"
	}

	return strings.Join([]string{"TagFilter", string(data)}, " ")
}
