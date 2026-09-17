package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRapidGrowthThresholdNewRequestBody 设置快速增长阈值请求体
type SetRapidGrowthThresholdNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 阈值
	Threshold int64 `json:"threshold"`
}

func (o SetRapidGrowthThresholdNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRapidGrowthThresholdNewRequestBody struct{}"
	}

	return strings.Join([]string{"SetRapidGrowthThresholdNewRequestBody", string(data)}, " ")
}
