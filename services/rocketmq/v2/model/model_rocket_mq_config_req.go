package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RocketMqConfigReq struct {

	// **参数解释**： RocketMQ配置名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name RocketMqConfigReqName `json:"name"`

	// **参数解释**： RocketMQ配置目标值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o RocketMqConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqConfigReq struct{}"
	}

	return strings.Join([]string{"RocketMqConfigReq", string(data)}, " ")
}

type RocketMqConfigReqName struct {
	value string
}

type RocketMqConfigReqNameEnum struct {
	FILE_RESERVED_TIME RocketMqConfigReqName
}

func GetRocketMqConfigReqNameEnum() RocketMqConfigReqNameEnum {
	return RocketMqConfigReqNameEnum{
		FILE_RESERVED_TIME: RocketMqConfigReqName{
			value: "fileReservedTime",
		},
	}
}

func (c RocketMqConfigReqName) Value() string {
	return c.value
}

func (c RocketMqConfigReqName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RocketMqConfigReqName) UnmarshalJSON(b []byte) error {
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
