package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IcebergSortOrder 排序顺序规格，决定了数据如何被重新排序以优化查询性能
type IcebergSortOrder struct {

	// 排序规范的id
	OrderId int32 `json:"order_id"`

	// IcebergSortField的数组
	Fields []IcebergSortField `json:"fields"`
}

func (o IcebergSortOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergSortOrder struct{}"
	}

	return strings.Join([]string{"IcebergSortOrder", string(data)}, " ")
}
