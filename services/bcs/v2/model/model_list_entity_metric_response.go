package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEntityMetricResponse struct {
	// 指标对象列表。

	Metrics        *[]EntityMetricList `json:"metrics,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEntityMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntityMetricResponse struct{}"
	}

	return strings.Join([]string{"ListEntityMetricResponse", string(data)}, " ")
}
