package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricThresholdResponse Response Object
type ShowMetricThresholdResponse struct {

	// 数据库类型
	EngineType *string `json:"engine_type,omitempty"`

	// 指标阈值列表
	Items          *[]MetricThresholdItem `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowMetricThresholdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricThresholdResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricThresholdResponse", string(data)}, " ")
}
