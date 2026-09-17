package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeMetricsResponse Response Object
type ShowNodeMetricsResponse struct {

	// 指标值
	Metrics        *[]MetricsInfo `json:"metrics,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowNodeMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeMetricsResponse", string(data)}, " ")
}
