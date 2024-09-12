package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecondaryIndexInfo struct {

	// 索引状态。 - 长度：[1, 255] - 取值字符限制：[a-z0-9_-]+
	IndexName string `bson:"index_name"`

	// 索引状态。 - \"creating\" - \"active\" - \"deleting\"
	IndexStatus string `bson:"index_status"`
}

func (o SecondaryIndexInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecondaryIndexInfo struct{}"
	}

	return strings.Join([]string{"SecondaryIndexInfo", string(data)}, " ")
}
