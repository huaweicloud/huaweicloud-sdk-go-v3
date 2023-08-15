package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionsRequest Request Object
type ShowPartitionsRequest struct {

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 表名
	TableName string `json:"table_name"`

	// 显示个数，默认值为100
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowPartitionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionsRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionsRequest", string(data)}, " ")
}
