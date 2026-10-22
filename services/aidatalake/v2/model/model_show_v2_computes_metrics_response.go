package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2ComputesMetricsResponse Response Object
type ShowV2ComputesMetricsResponse struct {
	Data *[]ComputeResourceMetrics `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowV2ComputesMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2ComputesMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowV2ComputesMetricsResponse", string(data)}, " ")
}
