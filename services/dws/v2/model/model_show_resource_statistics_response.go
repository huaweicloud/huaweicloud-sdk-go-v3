package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceStatisticsResponse Response Object
type ShowResourceStatisticsResponse struct {
	ClusterStatistics *StatusStatistics `json:"cluster_statistics,omitempty"`

	NodeStatistics *StatusStatistics `json:"node_statistics,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowResourceStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceStatisticsResponse", string(data)}, " ")
}
