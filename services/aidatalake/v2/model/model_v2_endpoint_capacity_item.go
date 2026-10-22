package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2EndpointCapacityItem **参数解释**：容量配置。
type V2EndpointCapacityItem struct {

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：   - CPU：中央处理器。   - GPU：图形处理器。   - NPU：神经网络处理器。 **默认取值**：不涉及。
	Type V2EndpointCapacityItemType `json:"type"`

	// **参数解释**：资源数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：1~10000。 **默认取值**：不涉及。
	MaximumCapacity int32 `json:"maximum_capacity"`

	// **参数解释**：最小资源数量。CPU资源单位为核，GPU/NPU资源单位为个。 **约束限制**：不涉及。 **取值范围**：0~10000，小于等于maximum_capacity。 **默认取值**：不涉及。
	MinimumCapacity int32 `json:"minimum_capacity"`

	// **参数解释**：资源规格。获取方式：从规格列表查询获取。 **约束限制**：预留资源和混合模式端点此参数必填。端点创建后此参数不支持变更。 **取值范围**：支持英文字母、数字、小数点及连字符，长度为1~64个字符。 **默认取值**：不涉及。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**：容量是否使用弹性资源。 **约束限制**：端点创建后此参数不支持变更。 **取值范围**：   - true：使用弹性资源。   - false：不使用弹性资源。 **默认取值**：false。
	UseOndemandResource *bool `json:"use_ondemand_resource,omitempty"`
}

func (o V2EndpointCapacityItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointCapacityItem struct{}"
	}

	return strings.Join([]string{"V2EndpointCapacityItem", string(data)}, " ")
}

type V2EndpointCapacityItemType struct {
	value string
}

type V2EndpointCapacityItemTypeEnum struct {
	CPU V2EndpointCapacityItemType
	GPU V2EndpointCapacityItemType
	NPU V2EndpointCapacityItemType
}

func GetV2EndpointCapacityItemTypeEnum() V2EndpointCapacityItemTypeEnum {
	return V2EndpointCapacityItemTypeEnum{
		CPU: V2EndpointCapacityItemType{
			value: "CPU",
		},
		GPU: V2EndpointCapacityItemType{
			value: "GPU",
		},
		NPU: V2EndpointCapacityItemType{
			value: "NPU",
		},
	}
}

func (c V2EndpointCapacityItemType) Value() string {
	return c.value
}

func (c V2EndpointCapacityItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2EndpointCapacityItemType) UnmarshalJSON(b []byte) error {
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
