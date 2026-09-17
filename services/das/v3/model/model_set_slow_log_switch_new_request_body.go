package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSlowLogSwitchNewRequestBody 设置慢日志开关请求体
type SetSlowLogSwitchNewRequestBody struct {

	// 是否开启
	SwitchOn bool `json:"switch_on"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 保存时长
	RetentionHours int64 `json:"retention_hours"`
}

func (o SetSlowLogSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSlowLogSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetSlowLogSwitchNewRequestBody", string(data)}, " ")
}
