package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlSwitchNewRequestBody 批量设置SQL开关请求体
type BatchSetSqlSwitchNewRequestBody struct {

	// 是否开启
	SwitchOn bool `json:"switch_on"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 设置开关的类型，取值范围：fullsql、slowsql
	SwitchType string `json:"switch_type"`

	// 实例ID列表
	InstanceIds []string `json:"instance_ids"`

	// 保存时长
	RetentionHours int64 `json:"retention_hours"`
}

func (o BatchSetSqlSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetSqlSwitchNewRequestBody", string(data)}, " ")
}
