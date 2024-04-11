package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancesStatisticsResponse Response Object
type ShowInstancesStatisticsResponse struct {

	// 实例总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例统计信息
	InstancesStatistics *[]InstancesStatisticsResponseBodyInstancesStatistics `json:"instances_statistics,omitempty"`
	HttpStatusCode      int                                                   `json:"-"`
}

func (o ShowInstancesStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancesStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstancesStatisticsResponse", string(data)}, " ")
}
