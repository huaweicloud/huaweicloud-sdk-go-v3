package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlobalSecondaryIndexInfo struct {

	// 二级索引名称。
	IndexName string `bson:"index_name"`

	// 二级索引名称。 - \"creating\" - \"active\" - \"deleting\"
	IndexStatus string `bson:"index_status"`
}

func (o GlobalSecondaryIndexInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalSecondaryIndexInfo struct{}"
	}

	return strings.Join([]string{"GlobalSecondaryIndexInfo", string(data)}, " ")
}
