package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineInstanceExtendProductInfoRequest Request Object
type ShowEngineInstanceExtendProductInfoRequest struct {

	// 消息引擎的类型。支持的类型为rabbitmq。
	Engine ShowEngineInstanceExtendProductInfoRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 产品的类型。 **约束限制**： 不涉及。 **取值范围**： advanced：专享版 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ShowEngineInstanceExtendProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoRequest", string(data)}, " ")
}

type ShowEngineInstanceExtendProductInfoRequestEngine struct {
	value string
}

type ShowEngineInstanceExtendProductInfoRequestEngineEnum struct {
	RABBITMQ ShowEngineInstanceExtendProductInfoRequestEngine
}

func GetShowEngineInstanceExtendProductInfoRequestEngineEnum() ShowEngineInstanceExtendProductInfoRequestEngineEnum {
	return ShowEngineInstanceExtendProductInfoRequestEngineEnum{
		RABBITMQ: ShowEngineInstanceExtendProductInfoRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoRequestEngine) UnmarshalJSON(b []byte) error {
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
