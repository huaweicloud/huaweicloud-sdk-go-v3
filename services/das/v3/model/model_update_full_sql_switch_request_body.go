package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFullSqlSwitchRequestBody 全量SQL开关请求体
type UpdateFullSqlSwitchRequestBody struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 开关。取值范围：1（开启）、0（关闭）
	OpenSwitch int32 `json:"open_switch"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 保存时长
	RetentionHours *int64 `json:"retention_hours,omitempty"`
}

func (o UpdateFullSqlSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFullSqlSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFullSqlSwitchRequestBody", string(data)}, " ")
}
