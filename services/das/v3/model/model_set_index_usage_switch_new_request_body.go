package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetIndexUsageSwitchNewRequestBody 设置索引使用开关请求体
type SetIndexUsageSwitchNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 开关状态
	SwitchOn bool `json:"switch_on"`
}

func (o SetIndexUsageSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetIndexUsageSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetIndexUsageSwitchNewRequestBody", string(data)}, " ")
}
