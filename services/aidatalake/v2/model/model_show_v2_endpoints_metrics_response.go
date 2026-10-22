package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2EndpointsMetricsResponse Response Object
type ShowV2EndpointsMetricsResponse struct {
	Data *[]V2EndpointMetrics `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowV2EndpointsMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2EndpointsMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowV2EndpointsMetricsResponse", string(data)}, " ")
}
