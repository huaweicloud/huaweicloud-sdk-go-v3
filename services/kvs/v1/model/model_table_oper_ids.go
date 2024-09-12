package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableOperIds 某个表的行操作数组， 可以是多个文档的put_kv或delete_kv。
type TableOperIds struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	KvOperIds *KvOperIds `bson:"kv_oper_ids"`
}

func (o TableOperIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableOperIds struct{}"
	}

	return strings.Join([]string{"TableOperIds", string(data)}, " ")
}
