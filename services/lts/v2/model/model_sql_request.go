package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlRequest struct {

	// 是时间范围相对
	IsTimeRangeRelative *bool `json:"is_time_range_relative,omitempty" xml:"is_time_range_relative"`

	// 日志流id
	LogStreamId string `json:"log_stream_id" xml:"log_stream_id"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty" xml:"log_stream_name"`

	// 日志组id
	LogGroupId string `json:"log_group_id" xml:"log_group_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty" xml:"log_group_name"`

	// sql语句
	Sql string `json:"sql" xml:"sql"`

	// 图表名称
	SqlRequestTitle string `json:"sql_request_title" xml:"sql_request_title"`

	// 查询执行任务时最近数据的时间范围(当search_time_range_unit为minute，则最大值为60;当search_time_range_unit为hour，则最大值为24)
	SearchTimeRange int32 `json:"search_time_range" xml:"search_time_range"`

	// 查询时间单位
	SearchTimeRangeUnit SqlRequestSearchTimeRangeUnit `json:"search_time_range_unit" xml:"search_time_range_unit"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
