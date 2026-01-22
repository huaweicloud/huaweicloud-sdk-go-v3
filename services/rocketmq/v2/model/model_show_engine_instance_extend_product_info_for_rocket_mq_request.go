package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEngineInstanceExtendProductInfoForRocketMqRequest Request Object
type ShowEngineInstanceExtendProductInfoForRocketMqRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 产品的类型。 advanced：专享版。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： advanced。
	Type *ShowEngineInstanceExtendProductInfoForRocketMqRequestType `json:"type,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询，offset大于等于0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowEngineInstanceExtendProductInfoForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoForRocketMqRequest", string(data)}, " ")
}

type ShowEngineInstanceExtendProductInfoForRocketMqRequestType struct {
	value string
}

type ShowEngineInstanceExtendProductInfoForRocketMqRequestTypeEnum struct {
	ADVANCED ShowEngineInstanceExtendProductInfoForRocketMqRequestType
}

func GetShowEngineInstanceExtendProductInfoForRocketMqRequestTypeEnum() ShowEngineInstanceExtendProductInfoForRocketMqRequestTypeEnum {
	return ShowEngineInstanceExtendProductInfoForRocketMqRequestTypeEnum{
		ADVANCED: ShowEngineInstanceExtendProductInfoForRocketMqRequestType{
			value: "advanced",
		},
	}
}

func (c ShowEngineInstanceExtendProductInfoForRocketMqRequestType) Value() string {
	return c.value
}

func (c ShowEngineInstanceExtendProductInfoForRocketMqRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEngineInstanceExtendProductInfoForRocketMqRequestType) UnmarshalJSON(b []byte) error {
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
