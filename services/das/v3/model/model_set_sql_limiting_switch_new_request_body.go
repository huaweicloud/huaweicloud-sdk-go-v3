package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlLimitingSwitchNewRequestBody 设置SQL限流开关请求体
type SetSqlLimitingSwitchNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 开关状态
	SwitchOn bool `json:"switch_on"`
}

func (o SetSqlLimitingSwitchNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlLimitingSwitchNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetSqlLimitingSwitchNewRequestBody", string(data)}, " ")
}
