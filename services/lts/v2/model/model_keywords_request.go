package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeywordsRequest struct {

	// 日志流id
	LogStreamId string `json:"log_stream_id"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 日志组id
	LogGroupId string `json:"log_group_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 关键词
	Keywords string `json:"keywords"`

	// 条件
	Condition KeywordsRequestCondition `json:"condition"`

	// 行数
	Number int32 `json:"number"`

	// 查询执行任务时最近数据的时间范围，最大值为60
	SearchTimeRange int32 `json:"search_time_range"`

	// 查询时间单位
	SearchTimeRangeUnit KeywordsRequestSearchTimeRangeUnit `json:"search_time_range_unit"`
}

func (o KeywordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsRequest struct{}"
	}

	return strings.Join([]string{"KeywordsRequest", string(data)}, " ")
}

type KeywordsRequestCondition struct {
	value string
}

type KeywordsRequestConditionEnum struct {
	GREATER_THAN_OR_EQUAL_TO KeywordsRequestCondition
	LESS_THAN_OR_EQUAL_TO    KeywordsRequestCondition
	LESS_THAN                KeywordsRequestCondition
	GREATER_THAN             KeywordsRequestCondition
}

func GetKeywordsRequestConditionEnum() KeywordsRequestConditionEnum {
	return KeywordsRequestConditionEnum{
		GREATER_THAN_OR_EQUAL_TO: KeywordsRequestCondition{
			value: ">=",
		},
		LESS_THAN_OR_EQUAL_TO: KeywordsRequestCondition{
			value: "<=",
		},
		LESS_THAN: KeywordsRequestCondition{
			value: "<",
		},
		GREATER_THAN: KeywordsRequestCondition{
			value: ">",
		},
	}
}

func (c KeywordsRequestCondition) Value() string {
	return c.value
}

func (c KeywordsRequestCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsRequestCondition) UnmarshalJSON(b []byte) error {
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

type KeywordsRequestSearchTimeRangeUnit struct {
	value string
}

type KeywordsRequestSearchTimeRangeUnitEnum struct {
	MINUTE KeywordsRequestSearchTimeRangeUnit
}

func GetKeywordsRequestSearchTimeRangeUnitEnum() KeywordsRequestSearchTimeRangeUnitEnum {
	return KeywordsRequestSearchTimeRangeUnitEnum{
		MINUTE: KeywordsRequestSearchTimeRangeUnit{
			value: "minute",
		},
	}
}

func (c KeywordsRequestSearchTimeRangeUnit) Value() string {
	return c.value
}

func (c KeywordsRequestSearchTimeRangeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsRequestSearchTimeRangeUnit) UnmarshalJSON(b []byte) error {
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
