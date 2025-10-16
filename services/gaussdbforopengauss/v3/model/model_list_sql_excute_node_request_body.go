package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListSqlExcuteNodeRequestBody struct {

	// **参数解释**: 类型。 **约束限制**: 不涉及。 **取值范围**: - slow  **默认取值**: 不涉及。
	Action ListSqlExcuteNodeRequestBodyAction `json:"action"`
}

func (o ListSqlExcuteNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlExcuteNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlExcuteNodeRequestBody", string(data)}, " ")
}

type ListSqlExcuteNodeRequestBodyAction struct {
	value string
}

type ListSqlExcuteNodeRequestBodyActionEnum struct {
	SLOW ListSqlExcuteNodeRequestBodyAction
}

func GetListSqlExcuteNodeRequestBodyActionEnum() ListSqlExcuteNodeRequestBodyActionEnum {
	return ListSqlExcuteNodeRequestBodyActionEnum{
		SLOW: ListSqlExcuteNodeRequestBodyAction{
			value: "slow",
		},
	}
}

func (c ListSqlExcuteNodeRequestBodyAction) Value() string {
	return c.value
}

func (c ListSqlExcuteNodeRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlExcuteNodeRequestBodyAction) UnmarshalJSON(b []byte) error {
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
