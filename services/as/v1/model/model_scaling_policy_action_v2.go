package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScalingPolicyActionV2 策略执行具体动作。
type ScalingPolicyActionV2 struct {

	// 操作选项，默认为ADD。 当scaling_resource_type为SCALING_GROUP，支持如下操作： - ADD：增加 - REMOVE/REDUCE：减少 - SET：设置为 当scaling_resource_type为BANDWIDTH，支持如下操作： - ADD：增加 - REDUCE：减少 - SET：设置为
	Operation *ScalingPolicyActionV2Operation `json:"operation,omitempty"`

	// 操作大小，取值范围为0到300的整数，默认为1。当scaling_resource_type为SCALING_GROUP时，size为实例个数,取值范围为0-300的整数，默认为1。当scaling_resource_type为BANDWIDTH时，size表示带宽大小，单位为Mbit/s，取值范围为1到300的整数，默认为1。当scaling_resource_type为SCALING_GROUP时，size和percentage参数只能选其中一个进行配置。
	Size *int32 `json:"size,omitempty"`

	// 操作百分比，取值为0到20000的整数。当scaling_resource_type为SCALING_GROUP时，size和instance_percentage参数均无配置，则size默认为1。当scaling_resource_type为BANDWIDTH时，不支持配置instance_percentage参数。
	Percentage *int32 `json:"percentage,omitempty"`

	// 操作限制。当scaling_resource_type为BANDWIDTH，且operation不为SET时，limits参数生效，单位为Mbit/s。此时，当operation为ADD时，limits表示带宽可调整的上限；当operation为REDUCE时，limits表示带宽可调整的下限。
	Limits *int32 `json:"limits,omitempty"`
}

func (o ScalingPolicyActionV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyActionV2 struct{}"
	}

	return strings.Join([]string{"ScalingPolicyActionV2", string(data)}, " ")
}

type ScalingPolicyActionV2Operation struct {
	value string
}

type ScalingPolicyActionV2OperationEnum struct {
	ADD    ScalingPolicyActionV2Operation
	REMOVE ScalingPolicyActionV2Operation
	REDUCE ScalingPolicyActionV2Operation
	SET    ScalingPolicyActionV2Operation
}

func GetScalingPolicyActionV2OperationEnum() ScalingPolicyActionV2OperationEnum {
	return ScalingPolicyActionV2OperationEnum{
		ADD: ScalingPolicyActionV2Operation{
			value: "ADD",
		},
		REMOVE: ScalingPolicyActionV2Operation{
			value: "REMOVE",
		},
		REDUCE: ScalingPolicyActionV2Operation{
			value: "REDUCE",
		},
		SET: ScalingPolicyActionV2Operation{
			value: "SET",
		},
	}
}

func (c ScalingPolicyActionV2Operation) Value() string {
	return c.value
}

func (c ScalingPolicyActionV2Operation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyActionV2Operation) UnmarshalJSON(b []byte) error {
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
