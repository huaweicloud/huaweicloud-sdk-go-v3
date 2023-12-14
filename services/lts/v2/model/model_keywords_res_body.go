package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type KeywordsResBody struct {

	// 日志流id
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 日志组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 关键词
	Keywords *string `json:"keywords,omitempty"`

	// 条件
	Condition *KeywordsResBodyCondition `json:"condition,omitempty"`

	// 行数
	Number *int32 `json:"number,omitempty"`

	// 查询执行任务时最近数据的时间范围，最大值为60
	SearchTimeRange *int32 `json:"search_time_range,omitempty"`

	// 查询时间单位
	SearchTimeRangeUnit *string `json:"search_time_range_unit,omitempty"`
}

func (o KeywordsResBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsResBody struct{}"
	}

	return strings.Join([]string{"KeywordsResBody", string(data)}, " ")
}

type KeywordsResBodyCondition struct {
	value string
}

type KeywordsResBodyConditionEnum struct {
	GREATER_THAN_OR_EQUAL_TO KeywordsResBodyCondition
	LESS_THAN_OR_EQUAL_TO    KeywordsResBodyCondition
	LESS_THAN                KeywordsResBodyCondition
	GREATER_THAN             KeywordsResBodyCondition
}

func GetKeywordsResBodyConditionEnum() KeywordsResBodyConditionEnum {
	return KeywordsResBodyConditionEnum{
		GREATER_THAN_OR_EQUAL_TO: KeywordsResBodyCondition{
			value: ">=",
		},
		LESS_THAN_OR_EQUAL_TO: KeywordsResBodyCondition{
			value: "<=",
		},
		LESS_THAN: KeywordsResBodyCondition{
			value: "<",
		},
		GREATER_THAN: KeywordsResBodyCondition{
			value: ">",
		},
	}
}

func (c KeywordsResBodyCondition) Value() string {
	return c.value
}

func (c KeywordsResBodyCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordsResBodyCondition) UnmarshalJSON(b []byte) error {
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
