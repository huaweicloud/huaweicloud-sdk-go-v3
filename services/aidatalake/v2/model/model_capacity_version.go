package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CapacityVersion **参数解释**：历史容量配置。 **约束限制**：不涉及。
type CapacityVersion struct {

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：   - CPU：中央处理器。   - GPU：图形处理器。   - NPU：神经网络处理器。 **默认取值**：不涉及。
	Type *CapacityVersionType `json:"type,omitempty"`

	// **参数解释**：资源数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：0~48000。 **默认取值**：不涉及。
	Capacity *int32 `json:"capacity,omitempty"`

	// **参数解释**：目标资源数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：0~48000。 **默认取值**：不涉及。
	TargetCapacity *int32 `json:"target_capacity,omitempty"`

	// **参数解释**：资源规格。 **约束限制**：不涉及。 **取值范围**：由英文字母、数字、小数点和中划线组成，长度为1~64个字符。 **默认取值**：不涉及。
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o CapacityVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityVersion struct{}"
	}

	return strings.Join([]string{"CapacityVersion", string(data)}, " ")
}

type CapacityVersionType struct {
	value string
}

type CapacityVersionTypeEnum struct {
	CPU CapacityVersionType
	GPU CapacityVersionType
	NPU CapacityVersionType
}

func GetCapacityVersionTypeEnum() CapacityVersionTypeEnum {
	return CapacityVersionTypeEnum{
		CPU: CapacityVersionType{
			value: "CPU",
		},
		GPU: CapacityVersionType{
			value: "GPU",
		},
		NPU: CapacityVersionType{
			value: "NPU",
		},
	}
}

func (c CapacityVersionType) Value() string {
	return c.value
}

func (c CapacityVersionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CapacityVersionType) UnmarshalJSON(b []byte) error {
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
