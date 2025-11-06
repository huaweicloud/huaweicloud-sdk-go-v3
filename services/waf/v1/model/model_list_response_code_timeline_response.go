package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResponseCodeTimelineResponse Response Object
type ListResponseCodeTimelineResponse struct {

	// **参数解释：** 安全统计的时间线，按时间顺序展示统计数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListResponseCodeTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponseCodeTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListResponseCodeTimelineResponse", string(data)}, " ")
}
