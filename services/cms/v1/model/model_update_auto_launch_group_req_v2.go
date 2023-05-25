package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type UpdateAutoLaunchGroupReqV2 struct {

	// 智能购买组名称(1-64个字符)，只能包含中文、字母、数字、下划线和中划线
	Name *string `json:"name,omitempty"`

	// 保障计划id
	GuaranteePlanId *string `json:"guarantee_plan_id,omitempty"`

	// 供应组目标容量，实例数量或者CPU个数，目标容量大于等于stable_capacity。spot实例的容量为满配容量减去stable_capacity。
	TargetCapacity *int32 `json:"target_capacity,omitempty"`

	// 按需实例目标容量(实例数量或CPU个数),小于target_capacity,供应组中可以没有按需实例。
	StableCapacity *int32 `json:"stable_capacity,omitempty"`

	// 超过目标容量时（目标容量减少）实例中断行为，枚举值 terminate：释放 noTermination：不释放 默认值：terminate
	ExcessFulfilledCapacityBehavior *UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior `json:"excess_fulfilled_capacity_behavior,omitempty"`

	// 请求到期正在运行实例中断行为，枚举值 terminate：释放 noTermination：不释放 默认值：terminate
	InstancesBehaviorWithExpiration *UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration `json:"instances_behavior_with_expiration,omitempty"`

	// 用户愿意为竞价实例每小时支付的最高价格。如果overrides中没有提供价格，可以使用该价格
	SpotPrice *float64 `json:"spot_price,omitempty"`
}

func (o UpdateAutoLaunchGroupReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoLaunchGroupReqV2 struct{}"
	}

	return strings.Join([]string{"UpdateAutoLaunchGroupReqV2", string(data)}, " ")
}

type UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior struct {
	value string
}

type UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum struct {
	TERMINATE      UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior
	NO_TERMINATION UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior
}

func GetUpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum() UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum {
	return UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum{
		TERMINATE: UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior{
			value: "terminate",
		},
		NO_TERMINATION: UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior{
			value: "noTermination",
		},
	}
}

func (c UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) Value() string {
	return c.value
}

func (c UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration struct {
	value string
}

type UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum struct {
	TERMINATE      UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration
	NO_TERMINATION UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration
}

func GetUpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum() UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum {
	return UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum{
		TERMINATE: UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration{
			value: "terminate",
		},
		NO_TERMINATION: UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration{
			value: "noTermination",
		},
	}
}

func (c UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration) Value() string {
	return c.value
}

func (c UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAutoLaunchGroupReqV2InstancesBehaviorWithExpiration) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
