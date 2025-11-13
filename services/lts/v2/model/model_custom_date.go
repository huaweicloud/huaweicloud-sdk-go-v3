package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomDate 告警规则配置的自定义查询区间。
type CustomDate struct {

	// **参数解释：** 查询起始时间相对于告警计划执行时间的偏移量数值。 **约束限制：** 根据start_unit字段的单位将有不同的取值范围，根据单位将start换算为时间不超过1天。 **取值范围：** start_unit取值为second，start取值范围为 1-86400； start_unit取值为minute，start取值范围为 1-1440； start_unit取值为hour，start取值范围为 1-24。 **默认取值：** -无。
	Start int32 `json:"start"`

	// **参数解释：** 查询起始时间相对于告警计划执行时间的偏移量单位。 **约束限制：** 整点时间即is_time_range_relative为false时start_unit不支持second。 **取值范围：** - hour - minute - second **默认取值：** minute
	StartUnit string `json:"start_unit"`

	// **参数解释：** 查询结束时间相对于告警计划执行时间的偏移量数值。 **约束限制：** 根据end_unit字段的单位将有不同的取值范围，根据单位将end换算为时间不超过1天； **取值范围：** end_unit取值为second，endt取值范围为 0-86399； end_unit取值为minute，end取值范围为 0-1439； end_unit取值为hour，end取值范围为 0-23。 **默认取值：** -无。
	End int32 `json:"end"`

	// **参数解释：** 查询结束时间相对于告警计划执行时间的偏移量单位。 **约束限制：** 整点时间即is_time_range_relative为false时end_unit不支持second。 **取值范围：** - hour - minute - second **默认取值：** minute
	EndUnit string `json:"end_unit"`
}

func (o CustomDate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomDate struct{}"
	}

	return strings.Join([]string{"CustomDate", string(data)}, " ")
}
