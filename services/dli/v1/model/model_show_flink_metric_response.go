package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkMetricResponse Response Object
type ShowFlinkMetricResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	Metrics        *FlinkMetric `json:"metrics,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowFlinkMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkMetricResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkMetricResponse", string(data)}, " ")
}
