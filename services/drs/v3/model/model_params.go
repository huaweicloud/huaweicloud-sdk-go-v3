package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据参数信息体
type Params struct {
	// 参数对比结果

	CompareResult *ParamsCompareResult `json:"compare_result,omitempty"`
	// 参数类型

	DataType *string `json:"data_type,omitempty"`
	// 分组

	Group *ParamsGroup `json:"group,omitempty"`
	// 参数名

	Key *string `json:"key,omitempty"`
	// 是否需要重启

	NeedRestart *ParamsNeedRestart `json:"need_restart,omitempty"`
	// 源数据库参数值

	SourceValue *string `json:"source_value,omitempty"`
	// 目标数据库参数值

	TargetValue *string `json:"target_value,omitempty"`
	// 参数范围

	ValueRange *string `json:"value_range,omitempty"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o Params) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Params struct{}"
	}

	return strings.Join([]string{"Params", string(data)}, " ")
}

type ParamsCompareResult struct {
	value string
}

type ParamsCompareResultEnum struct {
	TRUE  ParamsCompareResult
	FALSE ParamsCompareResult
}

func GetParamsCompareResultEnum() ParamsCompareResultEnum {
	return ParamsCompareResultEnum{
		TRUE: ParamsCompareResult{
			value: "true",
		},
		FALSE: ParamsCompareResult{
			value: "false",
		},
	}
}

func (c ParamsCompareResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsCompareResult) UnmarshalJSON(b []byte) error {
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

type ParamsGroup struct {
	value string
}

type ParamsGroupEnum struct {
	COMMON      ParamsGroup
	PERFORMANCE ParamsGroup
}

func GetParamsGroupEnum() ParamsGroupEnum {
	return ParamsGroupEnum{
		COMMON: ParamsGroup{
			value: "common-常规参数",
		},
		PERFORMANCE: ParamsGroup{
			value: "performance-性能参数",
		},
	}
}

func (c ParamsGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsGroup) UnmarshalJSON(b []byte) error {
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

type ParamsNeedRestart struct {
	value string
}

type ParamsNeedRestartEnum struct {
	TRUE  ParamsNeedRestart
	FALSE ParamsNeedRestart
}

func GetParamsNeedRestartEnum() ParamsNeedRestartEnum {
	return ParamsNeedRestartEnum{
		TRUE: ParamsNeedRestart{
			value: "true",
		},
		FALSE: ParamsNeedRestart{
			value: "false",
		},
	}
}

func (c ParamsNeedRestart) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParamsNeedRestart) UnmarshalJSON(b []byte) error {
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
