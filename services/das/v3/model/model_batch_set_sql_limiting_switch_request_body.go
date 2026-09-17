package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlLimitingSwitchRequestBody Sql Limitbatch Set Switch New请求体
type BatchSetSqlLimitingSwitchRequestBody struct {

	// 开关状态
	SwitchOn bool `json:"switch_on"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 设置开关的类型
	SwitchType string `json:"switch_type"`

	// 实例ID列表
	InstanceIds []string `json:"instance_ids"`
}

func (o BatchSetSqlLimitingSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlLimitingSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSetSqlLimitingSwitchRequestBody", string(data)}, " ")
}
