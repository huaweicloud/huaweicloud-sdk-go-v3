package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReturnedKvItemsOfTable struct {

	// 表名，仓内唯一。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName string `bson:"table_name"`

	ReturnedKvItems *[]ReturnedKvItem `bson:"returned_kv_items,omitempty"`
}

func (o ReturnedKvItemsOfTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReturnedKvItemsOfTable struct{}"
	}

	return strings.Join([]string{"ReturnedKvItemsOfTable", string(data)}, " ")
}
