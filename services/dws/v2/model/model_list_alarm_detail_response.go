package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmDetailResponse Response Object
type ListAlarmDetailResponse struct {

	// **参数解释**： 告警详情总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 告警列表。 **取值范围**： 不涉及。
	AlarmDetails   *[]AlarmDetailResponse `json:"alarm_details,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAlarmDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmDetailResponse", string(data)}, " ")
}
