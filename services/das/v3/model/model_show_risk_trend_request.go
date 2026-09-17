package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRiskTrendRequest Request Object
type ShowRiskTrendRequest struct {

	// 数据库类型
	EngineType string `json:"engine_type"`

	// 开始时间（Unix timestamp，毫秒）
	From int64 `json:"from"`

	// 结束时间（Unix timestamp，毫秒）
	To int64 `json:"to"`

	// 指标码
	MetricCode string `json:"metric_code"`
}

func (o ShowRiskTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowRiskTrendRequest", string(data)}, " ")
}
