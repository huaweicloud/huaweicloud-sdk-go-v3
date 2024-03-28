package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionsRequest Request Object
type ListPartitionsRequest struct {

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 表名
	TableName string `json:"table_name"`

	// 显示个数，默认值为100
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPartitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionsRequest struct{}"
	}

	return strings.Join([]string{"ListPartitionsRequest", string(data)}, " ")
}
