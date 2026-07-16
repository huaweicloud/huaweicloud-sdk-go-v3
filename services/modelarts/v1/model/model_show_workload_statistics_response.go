package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadStatisticsResponse Response Object
type ShowWorkloadStatisticsResponse struct {
	Statistics     *WorkloadListStatisticsStatistics `json:"statistics,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowWorkloadStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkloadStatisticsResponse", string(data)}, " ")
}
