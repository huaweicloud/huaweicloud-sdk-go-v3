package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMissingIndexSwitchNewRequestBody 设置缺失索引开关请求体
type SetMissingIndexSwitchNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 开关状态
	SwitchOn bool `json:"switch_on"`
}

func (o SetMissingIndexSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMissingIndexSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetMissingIndexSwitchNewRequestBody", string(data)}, " ")
}
