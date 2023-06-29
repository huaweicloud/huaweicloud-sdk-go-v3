package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiRequestPara struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 映射字段
	Mapping *string `json:"mapping,omitempty"`

	// 操作符
	Condition *ApiRequestParaCondition `json:"condition,omitempty"`
}

func (o ApiRequestPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiRequestPara struct{}"
	}

	return strings.Join([]string{"ApiRequestPara", string(data)}, " ")
}

type ApiRequestParaCondition struct {
	value string
}

type ApiRequestParaConditionEnum struct {
	CONDITION_TYPE_EQ     ApiRequestParaCondition
	CONDITION_TYPE_NE     ApiRequestParaCondition
	CONDITION_TYPE_GT     ApiRequestParaCondition
	CONDITION_TYPE_GE     ApiRequestParaCondition
	CONDITION_TYPE_LT     ApiRequestParaCondition
	CONDITION_TYPE_LE     ApiRequestParaCondition
	CONDITION_TYPE_LIKE   ApiRequestParaCondition
	CONDITION_TYPE_LIKE_L ApiRequestParaCondition
	CONDITION_TYPE_LIKE_R ApiRequestParaCondition
}

func GetApiRequestParaConditionEnum() ApiRequestParaConditionEnum {
	return ApiRequestParaConditionEnum{
		CONDITION_TYPE_EQ: ApiRequestParaCondition{
			value: "CONDITION_TYPE_EQ",
		},
		CONDITION_TYPE_NE: ApiRequestParaCondition{
			value: "CONDITION_TYPE_NE",
		},
		CONDITION_TYPE_GT: ApiRequestParaCondition{
			value: "CONDITION_TYPE_GT",
		},
		CONDITION_TYPE_GE: ApiRequestParaCondition{
			value: "CONDITION_TYPE_GE",
		},
		CONDITION_TYPE_LT: ApiRequestParaCondition{
			value: "CONDITION_TYPE_LT",
		},
		CONDITION_TYPE_LE: ApiRequestParaCondition{
			value: "CONDITION_TYPE_LE",
		},
		CONDITION_TYPE_LIKE: ApiRequestParaCondition{
			value: "CONDITION_TYPE_LIKE",
		},
		CONDITION_TYPE_LIKE_L: ApiRequestParaCondition{
			value: "CONDITION_TYPE_LIKE_L",
		},
		CONDITION_TYPE_LIKE_R: ApiRequestParaCondition{
			value: "CONDITION_TYPE_LIKE_R",
		},
	}
}

func (c ApiRequestParaCondition) Value() string {
	return c.value
}

func (c ApiRequestParaCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiRequestParaCondition) UnmarshalJSON(b []byte) error {
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
