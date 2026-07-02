package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ClientData struct {

	// **参数解释**： 客户端语言。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释**： 客户端版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 客户端ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClientId *string `json:"client_id,omitempty"`

	// **参数解释**： 客户端地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClientAddr *string `json:"client_addr,omitempty"`

	// **参数解释**： 消费类型。 **约束限制**： 不涉及。 **取值范围**： - CONSUME_POP：POP消费模式 - CONSUME_PASSIVELY：推消费模式 - CONSUME_ACTIVELY：拉消费模式 **默认取值**： 不涉及。
	ConsumeType *ClientDataConsumeType `json:"consume_type,omitempty"`

	// **参数解释**： 订阅关系列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}

func (o ClientData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientData struct{}"
	}

	return strings.Join([]string{"ClientData", string(data)}, " ")
}

type ClientDataConsumeType struct {
	value string
}

type ClientDataConsumeTypeEnum struct {
	CONSUME_POP       ClientDataConsumeType
	CONSUME_PASSIVELY ClientDataConsumeType
	CONSUME_ACTIVELY  ClientDataConsumeType
}

func GetClientDataConsumeTypeEnum() ClientDataConsumeTypeEnum {
	return ClientDataConsumeTypeEnum{
		CONSUME_POP: ClientDataConsumeType{
			value: "CONSUME_POP",
		},
		CONSUME_PASSIVELY: ClientDataConsumeType{
			value: "CONSUME_PASSIVELY",
		},
		CONSUME_ACTIVELY: ClientDataConsumeType{
			value: "CONSUME_ACTIVELY",
		},
	}
}

func (c ClientDataConsumeType) Value() string {
	return c.value
}

func (c ClientDataConsumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClientDataConsumeType) UnmarshalJSON(b []byte) error {
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
