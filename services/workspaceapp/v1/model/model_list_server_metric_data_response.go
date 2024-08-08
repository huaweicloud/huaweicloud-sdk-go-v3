package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerMetricDataResponse Response Object
type ListServerMetricDataResponse struct {

	// 监控数据。
	ServerMetrics  *[]ServerMetricData `json:"server_metrics,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListServerMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ListServerMetricDataResponse", string(data)}, " ")
}
