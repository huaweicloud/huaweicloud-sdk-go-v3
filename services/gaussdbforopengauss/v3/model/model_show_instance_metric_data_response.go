package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMetricDataResponse Response Object
type ShowInstanceMetricDataResponse struct {

	// **参数解释**: 实例ID。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 指标数据集合。
	Metrics        *[]MetricDataResult `json:"metrics,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowInstanceMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMetricDataResponse", string(data)}, " ")
}
