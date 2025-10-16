package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionWaitEventStatisticsResponse Response Object
type ListSessionWaitEventStatisticsResponse struct {

	// **参数解释**: 统计总条数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 统计数据。
	WaitEventInfo  *[]SessionWaitEventStatisticInfo `json:"wait_event_info,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListSessionWaitEventStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionWaitEventStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionWaitEventStatisticsResponse", string(data)}, " ")
}
