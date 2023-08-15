package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetricDataRequest Request Object
type CreateMetricDataRequest struct {

	// 添加一条或多条自定义指标监控数据，请求参数。
	Body *[]MetricDataItem `json:"body,omitempty"`
}

func (o CreateMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetricDataRequest struct{}"
	}

	return strings.Join([]string{"CreateMetricDataRequest", string(data)}, " ")
}
