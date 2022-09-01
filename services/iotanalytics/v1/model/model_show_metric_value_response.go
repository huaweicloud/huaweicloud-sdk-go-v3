package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMetricValueResponse struct {

	// 时间序列,使用UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss.SSS'Z',示例：2021-02-01T00:00:00.123Z
	Timestamps *[]string `json:"timestamps,omitempty" xml:"timestamps"`

	// 指标计算结果列表
	Metrics        *[]MetricValue `json:"metrics,omitempty" xml:"metrics"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowMetricValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricValueResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricValueResponse", string(data)}, " ")
}
