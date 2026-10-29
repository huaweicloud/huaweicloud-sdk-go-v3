package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeywordsRequestResponse struct {

	// **参数解释：** 日志流ID，获取方式请参见：[获取日志组ID和日志流ID](lts_api_0014.xml)。 **取值范围：** 不涉及。
	LogStreamId string `json:"log_stream_id"`

	// **参数解释：** 日志流名称。 **取值范围：** 不涉及。
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// **参数解释：** 日志流ID。 **取值范围：** 不涉及。
	LogGroupId string `json:"log_group_id"`

	// **参数解释：** 日志组名称。 **取值范围：** 不涉及。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// **参数解释：** 在日志搜索能查到的日志关键字。 **取值范围：** 不涉及。
	Keywords string `json:"keywords"`

	// **参数解释：** 告警查询条件。 **取值范围：**  <、- >、 <=、>=
	Condition string `json:"condition"`

	// **参数解释：** 告警匹配条数。 **取值范围：** 不涉及。
	Number int32 `json:"number"`

	// **参数解释：** 查询执行告警任务时最近数据的时间范围。 **取值范围：** - 最小值：1 - 最大值：60
	SearchTimeRange int32 `json:"search_time_range"`

	// **参数解释：** 查询告警时间范围单位。 **取值范围：** - minute - hour
	SearchTimeRangeUnit string `json:"search_time_range_unit"`

	CustomDate *CustomDate `json:"custom_date,omitempty"`

	// **参数解释：** 告警查询日志的时间区间为相对时间还是整点时间。 **约束限制：** 不涉及。 **取值范围：** - true: 相对时间。 - false: 整点时间。 **默认取值：** true
	IsTimeRangeRelative *bool `json:"is_time_range_relative,omitempty"`

	// **参数解释：** 支持根据不同条件触发不同级别的告警。当使用该参数时，将优先使用该参数中的条件触发告警，无视keywords_alarm_level参数以及keywords_requests中的condition、number参数。 **约束限制：** 不涉及。
	Conditions *[]KeywordsAlarmConditions `json:"conditions,omitempty"`

	// **参数解释：** 日志组别名。 **取值范围：** 不涉及。
	LogGroupNameAlias *string `json:"log_group_name_alias,omitempty"`

	// **参数解释：** 日志流别名。 **取值范围：** 不涉及。
	LogStreamNameAlias *string `json:"log_stream_name_alias,omitempty"`

	// **参数解释：** 企业项目ID。 **取值范围：** 不涉及。
	EpsId *string `json:"eps_id,omitempty"`

	TimeRange *TimeRange `json:"time_range,omitempty"`
}

func (o KeywordsRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsRequestResponse struct{}"
	}

	return strings.Join([]string{"KeywordsRequestResponse", string(data)}, " ")
}
