package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricItemsResponse Response Object
type ListMetricItemsResponse struct {
	MetaData *MetaData `json:"metaData,omitempty"`

	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListMetricItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricItemsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricItemsResponse", string(data)}, " ")
}
