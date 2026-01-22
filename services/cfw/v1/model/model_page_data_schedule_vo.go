package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageDataScheduleVo **参数解释**： 时间表响应信息 **取值范围**： 不涉及
type PageDataScheduleVo struct {

	// **参数解释**： 数量限制 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 时间表列表 **取值范围**： 不涉及
	Records *[]ScheduleVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o PageDataScheduleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageDataScheduleVo struct{}"
	}

	return strings.Join([]string{"PageDataScheduleVo", string(data)}, " ")
}
