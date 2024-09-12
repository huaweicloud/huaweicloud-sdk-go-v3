package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeTableResponse Response Object
type DescribeTableResponse struct {

	// 表名。 - 长度：[3, 63] - 取值字符限制：[a-z0-9_-]+
	TableName *string `bson:"table_name,omitempty"`

	PrimaryKeySchema *PrimaryKeySchema `bson:"primary_key_schema,omitempty"`

	// 本地二级索引模板，可以多个。
	LocalSecondaryIndexSchema *[]SecondaryIndex `bson:"local_secondary_index_schema,omitempty"`

	// 全局二级索引模板。
	GlobalSecondaryIndexSchema *[]GlobalSecondaryIndex `bson:"global_secondary_index_schema,omitempty"`

	RunTimeInfo    *RunTimeInfo `bson:"run_time_info,omitempty"`
	HttpStatusCode int          `bson:"-"`
}

func (o DescribeTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeTableResponse struct{}"
	}

	return strings.Join([]string{"DescribeTableResponse", string(data)}, " ")
}
