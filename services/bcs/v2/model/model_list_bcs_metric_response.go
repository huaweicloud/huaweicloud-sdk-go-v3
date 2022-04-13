package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBcsMetricResponse struct {
	// 指标对象列表。

	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListBcsMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsMetricResponse struct{}"
	}

	return strings.Join([]string{"ListBcsMetricResponse", string(data)}, " ")
}
