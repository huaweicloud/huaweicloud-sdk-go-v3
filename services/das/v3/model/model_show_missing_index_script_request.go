package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexScriptRequest Request Object
type ShowMissingIndexScriptRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 表名
	TableName string `json:"table_name"`

	// 相等列
	EqualityColumns string `json:"equality_columns"`

	// 不等列
	InequalityColumns string `json:"inequality_columns"`

	// 包含列
	IncludedColumns string `json:"included_columns"`

	// 对象ID
	ObjectId string `json:"object_id"`
}

func (o ShowMissingIndexScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexScriptRequest", string(data)}, " ")
}
