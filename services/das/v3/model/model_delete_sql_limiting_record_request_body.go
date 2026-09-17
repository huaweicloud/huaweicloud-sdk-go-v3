package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlLimitingRecordRequestBody 删除SQL限流记录请求体
type DeleteSqlLimitingRecordRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// SQL限流规则ID，可组合，用逗号分隔
	ItemIds string `json:"item_ids"`
}

func (o DeleteSqlLimitingRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitingRecordRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitingRecordRequestBody", string(data)}, " ")
}
