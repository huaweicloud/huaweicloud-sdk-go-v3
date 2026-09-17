package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DateInfo struct {

	// **参数解释**： 日期类型，如每月1号执行则为1th。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DateType *string `json:"date_type,omitempty"`

	// **参数解释**： 开始时间，如：04:00:00。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DateStart *string `json:"date_start,omitempty"`

	// **参数解释**： 结束时间，如：08:00:00。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DateEnd *string `json:"date_end,omitempty"`
}

func (o DateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DateInfo struct{}"
	}

	return strings.Join([]string{"DateInfo", string(data)}, " ")
}
