package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRisksRequest Request Object
type ListRisksRequest struct {

	// 开始时间（Unix timestamp，毫秒）
	From int64 `json:"from"`

	// 结束时间（Unix timestamp，毫秒）
	To int64 `json:"to"`

	// 数据库类型
	EngineType string `json:"engine_type"`

	// 返回TOP风险实例数量
	Num *int32 `json:"num,omitempty"`

	// 指标码
	MetricCode *string `json:"metric_code,omitempty"`
}

func (o ListRisksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRisksRequest struct{}"
	}

	return strings.Join([]string{"ListRisksRequest", string(data)}, " ")
}
