package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricsTrendResponse Response Object
type ListMetricsTrendResponse struct {

	// 查询指标趋势响应
	Body           *[]MetricsWithTime `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListMetricsTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricsTrendResponse struct{}"
	}

	return strings.Join([]string{"ListMetricsTrendResponse", string(data)}, " ")
}
