package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlRequest struct {

	// 日志流id
	LogStreamId string `json:"log_stream_id"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 日志组id
	LogGroupId string `json:"log_group_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// sql语句
	Sql string `json:"sql"`

	// 查询执行任务时最近数据的时间范围(当search_time_range_unit为minute，则最大值为60;当search_time_range_unit为hour，则最大值为24)
	SearchTimeRange *int32 `json:"search_time_range,omitempty"`

	// 查询时间单位
	SearchTimeRangeUnit *SqlRequestSearchTimeRangeUnit `json:"search_time_range_unit,omitempty"`

	CustomDate *CustomDate `json:"custom_date,omitempty"`

	// **参数解释：** 告警查询日志的时间区间为相对时间还是整点时间。（暂不开放，后续aom上线该功能后一起开放） **约束限制：** 不涉及。 **取值范围：** - true: 相对时间。 - false: 整点时间。 **默认取值：** true
	IsTimeRangeRelative *bool `json:"is_time_range_relative,omitempty"`
}

func (o SqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlRequest struct{}"
	}

	return strings.Join([]string{"SqlRequest", string(data)}, " ")
}

type SqlRequestSearchTimeRangeUnit struct {
	value string
}

type SqlRequestSearchTimeRangeUnitEnum struct {
	MINUTE SqlRequestSearchTimeRangeUnit
	HOUR   SqlRequestSearchTimeRangeUnit
}

func GetSqlRequestSearchTimeRangeUnitEnum() SqlRequestSearchTimeRangeUnitEnum {
	return SqlRequestSearchTimeRangeUnitEnum{
		MINUTE: SqlRequestSearchTimeRangeUnit{
			value: "minute",
		},
		HOUR: SqlRequestSearchTimeRangeUnit{
			value: "hour",
		},
	}
}

func (c SqlRequestSearchTimeRangeUnit) Value() string {
	return c.value
}

func (c SqlRequestSearchTimeRangeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlRequestSearchTimeRangeUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
