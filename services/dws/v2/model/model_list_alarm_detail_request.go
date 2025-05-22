package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmDetailRequest Request Object
type ListAlarmDetailRequest struct {

	// 时区
	TimeZone string `json:"time_zone"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 不限制。
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmDetailRequest", string(data)}, " ")
}
