package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScaleNodePoolOptions 节点池伸缩选项配置
type ScaleNodePoolOptions struct {

	// 扩容状态检查策略: instant(同步检查), async(异步检查)。默认同步检查instant
	ScalableChecking *string `json:"scalableChecking,omitempty"`

	// **参数解释**： 扩容的策略，允许为空，该参数scaleGroups传多项时有效。 **约束限制**： 不涉及 **取值范围**： - AZBalance：AZ优先策略，扩容节点池时，系统会使各个AZ间的节点数尽可能的均衡，规格会在所选伸缩组中随机指定。该策略适用于对节点成本和可用区无特殊要求的场景，优点是配置简便、降低单点故障风险。注意：如果某个AZ资源不足，该AZ期望的扩容节点会向其他AZ扩容，可能会使AZ间节点不均衡。如需解决该问题，可在该AZ资源充足时尝试再次扩容。 - Random：随机策略，从下发的规格scaleGroups列表中随机选择伸缩组扩容。  **默认取值**： Random
	ScalePolicy *ScaleNodePoolOptionsScalePolicy `json:"scalePolicy,omitempty"`

	BillingConfigOverride *ScaleUpBillingConfigOverride `json:"billingConfigOverride,omitempty"`
}

func (o ScaleNodePoolOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleNodePoolOptions struct{}"
	}

	return strings.Join([]string{"ScaleNodePoolOptions", string(data)}, " ")
}

type ScaleNodePoolOptionsScalePolicy struct {
	value string
}

type ScaleNodePoolOptionsScalePolicyEnum struct {
	AZ_BALANCE ScaleNodePoolOptionsScalePolicy
	RANDOM     ScaleNodePoolOptionsScalePolicy
}

func GetScaleNodePoolOptionsScalePolicyEnum() ScaleNodePoolOptionsScalePolicyEnum {
	return ScaleNodePoolOptionsScalePolicyEnum{
		AZ_BALANCE: ScaleNodePoolOptionsScalePolicy{
			value: "AZBalance",
		},
		RANDOM: ScaleNodePoolOptionsScalePolicy{
			value: "Random",
		},
	}
}

func (c ScaleNodePoolOptionsScalePolicy) Value() string {
	return c.value
}

func (c ScaleNodePoolOptionsScalePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScaleNodePoolOptionsScalePolicy) UnmarshalJSON(b []byte) error {
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
