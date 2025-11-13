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
}

func (o KeywordsRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsRequestResponse struct{}"
	}

	return strings.Join([]string{"KeywordsRequestResponse", string(data)}, " ")
}
