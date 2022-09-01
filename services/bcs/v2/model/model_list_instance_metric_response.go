package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstanceMetricResponse struct {

	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty" xml:"metrics"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListInstanceMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMetricResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricResponse", string(data)}, " ")
}
