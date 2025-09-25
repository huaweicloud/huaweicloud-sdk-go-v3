package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionStatisticsResponse Response Object
type ListSessionStatisticsResponse struct {

	// **参数解释**: 总数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 会话统计列表。
	StatisticsList *[]DimensionListResult `json:"statistics_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSessionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionStatisticsResponse", string(data)}, " ")
}
