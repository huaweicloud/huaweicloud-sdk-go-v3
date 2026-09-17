package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSqlLimitingRecordNewRequestBody 更新SQL限流记录请求体
type UpdateSqlLimitingRecordNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// SQL限流规则ID，可组合，用逗号分隔
	ItemIds *string `json:"item_ids,omitempty"`

	// 最大并发数
	MaxConnection *int32 `json:"max_connection,omitempty"`

	// 最大等待时间
	MaxWaiting *int32 `json:"max_waiting,omitempty"`
}

func (o UpdateSqlLimitingRecordNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitingRecordNewRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitingRecordNewRequestBody", string(data)}, " ")
}
