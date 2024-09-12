package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DescribeTableRequestBody struct {

	// 表名。 - 长度：[3, 63] - 取值字符限制：[a-zA-Z0-9_-]+
	TableName string `bson:"table_name"`
}

func (o DescribeTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeTableRequestBody struct{}"
	}

	return strings.Join([]string{"DescribeTableRequestBody", string(data)}, " ")
}
