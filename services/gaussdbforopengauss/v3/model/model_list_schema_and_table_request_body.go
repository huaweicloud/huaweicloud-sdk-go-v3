package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemaAndTableRequestBody 查询table和schema请求体
type ListSchemaAndTableRequestBody struct {

	// SQL文本
	SqlText string `json:"sql_text"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListSchemaAndTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemaAndTableRequestBody struct{}"
	}

	return strings.Join([]string{"ListSchemaAndTableRequestBody", string(data)}, " ")
}
