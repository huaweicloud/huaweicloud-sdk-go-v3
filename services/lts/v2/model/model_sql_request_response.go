package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlRequestResponse struct {

	// **参数解释：** 日志流ID。 **取值范围：** 不涉及。
	LogStreamId string `json:"log_stream_id"`

	// **参数解释：** 日志流名称。 **取值范围：** 不涉及。
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// **参数解释：** 日志流ID。 **取值范围：** 不涉及。
	LogGroupId string `json:"log_group_id"`

	// **参数解释：** 日志组名称。 **取值范围：** 不涉及。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// **参数解释：** SQL语句。 **取值范围：** 不涉及。
	Sql string `json:"sql"`

	// **参数解释：** 配置告警关联的图表名称。 **取值范围：** 不涉及。
	SqlRequestTitle string `json:"sql_request_title"`

	// **参数解释：** 查询执行告警任务时最近数据的时间范围。 **取值范围：** 不涉及。
	SearchTimeRange int32 `json:"search_time_range"`

	// **参数解释：** 查询告警时间范围单位。 **取值范围：** - minute - hour
	SearchTimeRangeUnit string `json:"search_time_range_unit"`

	CustomDate *CustomDate `json:"custom_date,omitempty"`

	// **参数解释：** 告警查询日志的时间区间为相对时间还是整点时间。 **约束限制：** 不涉及。 **取值范围：** - true: 相对时间。 - false: 整点时间。 **默认取值：** true
	IsTimeRangeRelative *bool `json:"is_time_range_relative,omitempty"`
}

func (o SqlRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlRequestResponse struct{}"
	}

	return strings.Join([]string{"SqlRequestResponse", string(data)}, " ")
}
