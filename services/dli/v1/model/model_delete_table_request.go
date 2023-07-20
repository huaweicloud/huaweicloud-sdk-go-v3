package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableRequest Request Object
type DeleteTableRequest struct {

	// 待删除的表所在的数据库名称。
	DatabaseName string `json:"database_name"`

	// 待删除的表名称。
	TableName string `json:"table_name"`

	// 是否异步执行
	Async *bool `json:"async,omitempty"`
}

func (o DeleteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteTableRequest", string(data)}, " ")
}
