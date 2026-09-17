package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmDetailRequest Request Object
type ListAlarmDetailRequest struct {

	// **参数解释**： 时区信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： GMT+08:00
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 开始UTC时间，单位精确到毫秒。 **取值范围**： 不涉及。
	From *int64 `json:"from,omitempty"`

	// **参数解释**： 结束UTC时间，单位精确到毫秒。 **取值范围**： 不涉及。
	To *int64 `json:"to,omitempty"`

	// **参数解释**： 查询前N天到当前时间范围内的告警。 **约束限制**： 不涉及。 **取值范围**： 大于0。
	RecentDay *int32 `json:"recent_day,omitempty"`
}

func (o ListAlarmDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmDetailRequest", string(data)}, " ")
}
