package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 排列顺序
type Order struct {

	// 列的名称
	Column *string `json:"column,omitempty"`

	// 指示是按升序 (== 1) 还是降序 (==0) 对列进行排序
	SortOrder *int32 `json:"sort_order,omitempty"`
}

func (o Order) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Order struct{}"
	}

	return strings.Join([]string{"Order", string(data)}, " ")
}
