package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// This is a auto create Body Object
type AutoLaunchGroupReqV2 struct {

	// 智能购买组名称。 取值范围：1-64个字符，只能包含中文、字母、数字、下划线和中划线。
	Name string `json:"name"`

	// 创建智能购买组参数核查 true：发送检查请求，不会创建智能购买组。检查项包括是否填写了必需参数、请求格式等。 如果检查不通过，则返回对应错误。 如果检查通过，则返回202状态码。 false：发送正常请求，通过检查后并且执行创建智能购买组请求。
	DryRun *bool `json:"dry_run,omitempty"`

	// 请求类型，枚举值 request：一次性。仅在启动时交付实例集群，调度失败后不再重试。 maintain：持续供应。在启动时尝试交付实例集群，并监控实时容量，未达到目标容量则尝试继续创建ECS实例。 默认值：maintain
	Type *AutoLaunchGroupReqV2Type `json:"type,omitempty"`

	// 算力保障计划ID
	GuaranteePlanId *string `json:"guarantee_plan_id,omitempty"`

	// 智能购买组目标容量。 实例数量或者CPU个数目标容量大于等于stable_capacity。竞价实例的容量为满配容量减去stable_capacity。
	TargetCapacity int32 `json:"target_capacity"`

	// 按需实例目标容量。 目标容量指实例数量或CPU个数，必须小于等于target_capacity，智能购买组中可以没有按需实例。
	StableCapacity *int32 `json:"stable_capacity,omitempty"`

	// 超过目标容量或目标容量减少时的实例中断行为，枚举值 terminate：释放 noTermination：不释放 默认值：terminate
	ExcessFulfilledCapacityBehavior *AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior `json:"excess_fulfilled_capacity_behavior,omitempty"`

	// 请求到期正在的实例中断行为，枚举值 terminate：释放 noTermination：不释放 默认值：terminate
	InstancesBehaviorWithExpiration *AutoLaunchGroupReqV2InstancesBehaviorWithExpiration `json:"instances_behavior_with_expiration,omitempty"`

	// 请求开始时间，和valid_until共同确定有效时段。 按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ。 默认值：立即生效
	ValidSince *sdktime.SdkTime `json:"valid_since,omitempty"`

	// 请求结束时间，和valid_since共同确定有效时段。 按照ISO8601标准表示，并使用UTC +0时间，格式为yyyy-MM-ddTHH:mm:ssZ。 默认值：无限期
	ValidUntil *sdktime.SdkTime `json:"valid_until,omitempty"`

	// 实例分配策略，枚举值 lowest_price：价格最低策略，智能购买组购买的所有实例的价格总和最低。 prioritized：优先级策略，按照规格设定的优先级创建实例。  capacity_optimized：容量最优化策略，智能购买组购买的实例按照大规格优先进行购买。 默认值：lowest_price
	AllocationStrategy *AutoLaunchGroupReqV2AllocationStrategy `json:"allocation_strategy,omitempty"`

	// 智能购买组内各区域的资源描述
	RegionSpecs []RegionSpec `json:"region_specs"`

	// 资源供给中规格选择策略：枚举值 singlation：选择一种规格供给 multiple：组合多种规格供给 默认值：multiple
	SupplyOption *AutoLaunchGroupReqV2SupplyOption `json:"supply_option,omitempty"`

	// 用户愿意为竞价实例每小时支付的最高价格。如果overrides中没有提供价格，可以使用该价格
	SpotPrice *float64 `json:"spot_price,omitempty"`
}

func (o AutoLaunchGroupReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLaunchGroupReqV2 struct{}"
	}

	return strings.Join([]string{"AutoLaunchGroupReqV2", string(data)}, " ")
}

type AutoLaunchGroupReqV2Type struct {
	value string
}

type AutoLaunchGroupReqV2TypeEnum struct {
	REQUEST  AutoLaunchGroupReqV2Type
	MAINTAIN AutoLaunchGroupReqV2Type
}

func GetAutoLaunchGroupReqV2TypeEnum() AutoLaunchGroupReqV2TypeEnum {
	return AutoLaunchGroupReqV2TypeEnum{
		REQUEST: AutoLaunchGroupReqV2Type{
			value: "request",
		},
		MAINTAIN: AutoLaunchGroupReqV2Type{
			value: "maintain",
		},
	}
}

func (c AutoLaunchGroupReqV2Type) Value() string {
	return c.value
}

func (c AutoLaunchGroupReqV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLaunchGroupReqV2Type) UnmarshalJSON(b []byte) error {
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

type AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior struct {
	value string
}

type AutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum struct {
	TERMINATE      AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior
	NO_TERMINATION AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior
}

func GetAutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum() AutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum {
	return AutoLaunchGroupReqV2ExcessFulfilledCapacityBehaviorEnum{
		TERMINATE: AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior{
			value: "terminate",
		},
		NO_TERMINATION: AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior{
			value: "noTermination",
		},
	}
}

func (c AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) Value() string {
	return c.value
}

func (c AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLaunchGroupReqV2ExcessFulfilledCapacityBehavior) UnmarshalJSON(b []byte) error {
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

type AutoLaunchGroupReqV2InstancesBehaviorWithExpiration struct {
	value string
}

type AutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum struct {
	TERMINATE      AutoLaunchGroupReqV2InstancesBehaviorWithExpiration
	NO_TERMINATION AutoLaunchGroupReqV2InstancesBehaviorWithExpiration
}

func GetAutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum() AutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum {
	return AutoLaunchGroupReqV2InstancesBehaviorWithExpirationEnum{
		TERMINATE: AutoLaunchGroupReqV2InstancesBehaviorWithExpiration{
			value: "terminate",
		},
		NO_TERMINATION: AutoLaunchGroupReqV2InstancesBehaviorWithExpiration{
			value: "noTermination",
		},
	}
}

func (c AutoLaunchGroupReqV2InstancesBehaviorWithExpiration) Value() string {
	return c.value
}

func (c AutoLaunchGroupReqV2InstancesBehaviorWithExpiration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLaunchGroupReqV2InstancesBehaviorWithExpiration) UnmarshalJSON(b []byte) error {
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

type AutoLaunchGroupReqV2AllocationStrategy struct {
	value string
}

type AutoLaunchGroupReqV2AllocationStrategyEnum struct {
	LOWEST_PRICE       AutoLaunchGroupReqV2AllocationStrategy
	DIVERSIFIED        AutoLaunchGroupReqV2AllocationStrategy
	CAPACITY_OPTIMIZED AutoLaunchGroupReqV2AllocationStrategy
	PRIORITIZED        AutoLaunchGroupReqV2AllocationStrategy
}

func GetAutoLaunchGroupReqV2AllocationStrategyEnum() AutoLaunchGroupReqV2AllocationStrategyEnum {
	return AutoLaunchGroupReqV2AllocationStrategyEnum{
		LOWEST_PRICE: AutoLaunchGroupReqV2AllocationStrategy{
			value: "lowest_price",
		},
		DIVERSIFIED: AutoLaunchGroupReqV2AllocationStrategy{
			value: "diversified",
		},
		CAPACITY_OPTIMIZED: AutoLaunchGroupReqV2AllocationStrategy{
			value: "capacity_optimized",
		},
		PRIORITIZED: AutoLaunchGroupReqV2AllocationStrategy{
			value: "prioritized",
		},
	}
}

func (c AutoLaunchGroupReqV2AllocationStrategy) Value() string {
	return c.value
}

func (c AutoLaunchGroupReqV2AllocationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLaunchGroupReqV2AllocationStrategy) UnmarshalJSON(b []byte) error {
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

type AutoLaunchGroupReqV2SupplyOption struct {
	value string
}

type AutoLaunchGroupReqV2SupplyOptionEnum struct {
	SINGLATION AutoLaunchGroupReqV2SupplyOption
	MULTIPLE   AutoLaunchGroupReqV2SupplyOption
}

func GetAutoLaunchGroupReqV2SupplyOptionEnum() AutoLaunchGroupReqV2SupplyOptionEnum {
	return AutoLaunchGroupReqV2SupplyOptionEnum{
		SINGLATION: AutoLaunchGroupReqV2SupplyOption{
			value: "singlation",
		},
		MULTIPLE: AutoLaunchGroupReqV2SupplyOption{
			value: "multiple",
		},
	}
}

func (c AutoLaunchGroupReqV2SupplyOption) Value() string {
	return c.value
}

func (c AutoLaunchGroupReqV2SupplyOption) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoLaunchGroupReqV2SupplyOption) UnmarshalJSON(b []byte) error {
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
