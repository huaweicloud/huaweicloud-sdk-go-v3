package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableBatch 单个表的所有kv操作。
type TableBatch struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// kv操作数组。
	KvOpers []OperItem `bson:"kv_opers"`
}

func (o TableBatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableBatch struct{}"
	}

	return strings.Join([]string{"TableBatch", string(data)}, " ")
}
