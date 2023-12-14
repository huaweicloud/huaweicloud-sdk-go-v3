package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionMetricsResponse Response Object
type ShowPartitionMetricsResponse struct {
	Metrics        *Metrics `json:"metrics,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowPartitionMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowPartitionMetricsResponse", string(data)}, " ")
}
