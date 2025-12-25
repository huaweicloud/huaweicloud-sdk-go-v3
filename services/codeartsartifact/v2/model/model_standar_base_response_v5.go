package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StandarBaseResponseV5 struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *StandarBaseResponseV5Status `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`
}

func (o StandarBaseResponseV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandarBaseResponseV5 struct{}"
	}

	return strings.Join([]string{"StandarBaseResponseV5", string(data)}, " ")
}

type StandarBaseResponseV5Status struct {
	value string
}

type StandarBaseResponseV5StatusEnum struct {
	SUCCESS StandarBaseResponseV5Status
	ERROR   StandarBaseResponseV5Status
}

func GetStandarBaseResponseV5StatusEnum() StandarBaseResponseV5StatusEnum {
	return StandarBaseResponseV5StatusEnum{
		SUCCESS: StandarBaseResponseV5Status{
			value: "success",
		},
		ERROR: StandarBaseResponseV5Status{
			value: "error",
		},
	}
}

func (c StandarBaseResponseV5Status) Value() string {
	return c.value
}

func (c StandarBaseResponseV5Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StandarBaseResponseV5Status) UnmarshalJSON(b []byte) error {
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
