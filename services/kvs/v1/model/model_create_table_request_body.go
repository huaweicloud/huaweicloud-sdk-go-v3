package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTableRequestBody struct {

	// 表名，仓内唯一。
	TableName string `bson:"table_name"`

	// 表计费模式，可为\"provisioned\"或\"on_demand\" - 预置模式：provisioned - 按需模式：on_demand
	BillMode *string `bson:"bill_mode,omitempty"`

	ProvisionedThroughput *ProvisionedThroughput `bson:"provisioned_throughput,omitempty"`

	PrimaryKeySchema *PrimaryKeySchema `bson:"primary_key_schema"`

	// 本地二级索引模板，可以多个。
	LocalSecondaryIndexSchema *[]SecondaryIndex `bson:"local_secondary_index_schema,omitempty"`

	// 全局二级索引模板。
	GlobalSecondaryIndexSchema *[]GlobalSecondaryIndex `bson:"global_secondary_index_schema,omitempty"`

	PreSplitKeyOptions *PreSplitKeyOptions `bson:"pre_split_key_options,omitempty"`
}

func (o CreateTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTableRequestBody", string(data)}, " ")
}
