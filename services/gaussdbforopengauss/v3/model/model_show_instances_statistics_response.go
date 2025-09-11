package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancesStatisticsResponse Response Object
type ShowInstancesStatisticsResponse struct {

	// **参数解释**: 实例总数。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 实例统计信息。 **取值范围**: 不涉及。
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
