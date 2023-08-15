package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnableAction 停用启用企业项目操作
type EnableAction struct {

	// 启用操作
	Action EnableActionAction `json:"action"`
}

func (o EnableAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAction struct{}"
	}

	return strings.Join([]string{"EnableAction", string(data)}, " ")
}

type EnableActionAction struct {
	value string
}

type EnableActionActionEnum struct {
	ENABLE EnableActionAction
}

func GetEnableActionActionEnum() EnableActionActionEnum {
	return EnableActionActionEnum{
		ENABLE: EnableActionAction{
			value: "enable",
		},
	}
}

func (c EnableActionAction) Value() string {
	return c.value
}

func (c EnableActionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnableActionAction) UnmarshalJSON(b []byte) error {
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
