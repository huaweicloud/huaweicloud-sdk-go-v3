package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CapacityItem **参数解释**：最大容量配置。 **约束限制**：不涉及。
type CapacityItem struct {

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：   - CPU：中央处理器。   - GPU：图形处理器。   - NPU：神经网络处理器。 **默认取值**：不涉及。
	Type CapacityItemType `json:"type"`

	// **参数解释**：资源数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：0~48000。 **默认取值**：不涉及。
	Capacity int32 `json:"capacity"`

	// **参数解释**：资源规格。 **约束限制**：不涉及。 **取值范围**：由英文字母、数字、小数点和中划线组成，长度为1~64个字符。 **默认取值**：不涉及。
	SpecCode string `json:"spec_code"`
}

func (o CapacityItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityItem struct{}"
	}

	return strings.Join([]string{"CapacityItem", string(data)}, " ")
}

type CapacityItemType struct {
	value string
}

type CapacityItemTypeEnum struct {
	CPU CapacityItemType
	GPU CapacityItemType
	NPU CapacityItemType
}

func GetCapacityItemTypeEnum() CapacityItemTypeEnum {
	return CapacityItemTypeEnum{
		CPU: CapacityItemType{
			value: "CPU",
		},
		GPU: CapacityItemType{
			value: "GPU",
		},
		NPU: CapacityItemType{
			value: "NPU",
		},
	}
}

func (c CapacityItemType) Value() string {
	return c.value
}

func (c CapacityItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CapacityItemType) UnmarshalJSON(b []byte) error {
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
