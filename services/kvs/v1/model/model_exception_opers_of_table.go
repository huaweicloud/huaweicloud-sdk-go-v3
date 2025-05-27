package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExceptionOpersOfTable struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// 未处理的操作列表。 - 数组元素：未处理的操作标识。
	UnprocessedOpers *[]int32 `bson:"unprocessed_opers,omitempty"`

	// 失败的操作列表，可以是多个。
	FailedOpers *[]Fail `bson:"failed_opers,omitempty"`
}

func (o ExceptionOpersOfTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptionOpersOfTable struct{}"
	}

	return strings.Join([]string{"ExceptionOpersOfTable", string(data)}, " ")
}
