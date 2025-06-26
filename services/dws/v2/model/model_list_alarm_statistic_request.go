package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmStatisticRequest Request Object
type ListAlarmStatisticRequest struct {

	// **参数解释**： 时区信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： GMT+08:00
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ListAlarmStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmStatisticRequest", string(data)}, " ")
}
