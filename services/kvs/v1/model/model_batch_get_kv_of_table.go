package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetKvOfTable 单个表的所有kv操作。
type BatchGetKvOfTable struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	// 行操作数组，行操作类型可以是多个文档的get。
	GetKvOpers []GetOper `bson:"get_kv_opers"`
}

func (o BatchGetKvOfTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetKvOfTable struct{}"
	}

	return strings.Join([]string{"BatchGetKvOfTable", string(data)}, " ")
}
