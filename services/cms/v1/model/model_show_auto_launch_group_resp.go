package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowAutoLaunchGroupResp 智能购买组信息
type ShowAutoLaunchGroupResp struct {

	// 智能购买组名称
	Name string `json:"name"`

	// 请求类型。枚举值 request：一次性。仅在启动时交付实例集群，调度失败后不再重试。 maintain：持续供应。在启动时尝试交付实例集群，并监控实时容量，未达到目标容量则尝试继续创建ECS实例
	Type ShowAutoLaunchGroupRespType `json:"type"`

	// 智能购买组的运行状态，枚举值 SUBMITTED：已提交 ACTIVE：运行中 DELETING：删除中 DELETED：已删除
	Status string `json:"status"`

	// 智能购买组的任务状态，枚举值 HANDLING：购买中 FULFILLED：智能购买组已满配 ERROR：智能购买组异常
	TaskState string `json:"task_state"`

	// 智能购买组目标容量. 实例数量或者CPU个数，目标容量大于等于stable_capacity。竞价实例的容量为满配容量减去stable_capacity
	TargetCapacity int32 `json:"target_capacity"`

	// 按需实例目标容量 目标容量指实例数量或CPU个数，必须小于等于target_capacity，智能购买组中可以没有按需实例
	StableCapacity int32 `json:"stable_capacity"`

	// 当前已经购买成功的总算力
	CurrentCapacity int32 `json:"current_capacity"`

	// 当前已经购买成功的按需算力
	CurrentStableCapacity int32 `json:"current_stable_capacity"`

	// 超过目标容量或目标容量减少时的实例中断行为，枚举值 terminate：释放 noTermination：不释放
	ExcessFulfilledCapacityBehavior ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior `json:"excess_fulfilled_capacity_behavior"`

	// 请求到期时的实例中断行为，枚举值 terminate：释放 noTermination：不释放
	InstancesBehaviorWithExpiration ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration `json:"instances_behavior_with_expiration"`

	// 请求开始时间，格式为yyyy-MM-ddTHH:mm:ssZ
	ValidSince *sdktime.SdkTime `json:"valid_since"`

	// 请求结束时间，格式为yyyy-MM-ddTHH:mm:ssZ
	ValidUntil *sdktime.SdkTime `json:"valid_until"`

	// 智能购买组在各个区域的配置
	RegionSpecs []RegionSpec `json:"region_specs"`

	// 实例分配策略，枚举值 lowest_price：价格最低策略，智能购买组购买的所有实例的价格总和最低。 prioritized：优先级策略，按照规格设定的优先级创建实例。 capacity_optimized：容量最优化策略，智能购买组购买的实例按照大规格优先进行购买。
	AllocationStrategy ShowAutoLaunchGroupRespAllocationStrategy `json:"allocation_strategy"`

	// 用户愿意为竞价实例每小时支付的最高价格。如果overrides中没有提供价格，可以使用该价格。
	SpotPrice *float64 `json:"spot_price,omitempty"`

	// 智能购买组创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o ShowAutoLaunchGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoLaunchGroupResp struct{}"
	}

	return strings.Join([]string{"ShowAutoLaunchGroupResp", string(data)}, " ")
}

type ShowAutoLaunchGroupRespType struct {
	value string
}

type ShowAutoLaunchGroupRespTypeEnum struct {
	REQUEST  ShowAutoLaunchGroupRespType
	MAINTAIN ShowAutoLaunchGroupRespType
}

func GetShowAutoLaunchGroupRespTypeEnum() ShowAutoLaunchGroupRespTypeEnum {
	return ShowAutoLaunchGroupRespTypeEnum{
		REQUEST: ShowAutoLaunchGroupRespType{
			value: "request",
		},
		MAINTAIN: ShowAutoLaunchGroupRespType{
			value: "maintain",
		},
	}
}

func (c ShowAutoLaunchGroupRespType) Value() string {
	return c.value
}

func (c ShowAutoLaunchGroupRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoLaunchGroupRespType) UnmarshalJSON(b []byte) error {
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

type ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior struct {
	value string
}

type ShowAutoLaunchGroupRespExcessFulfilledCapacityBehaviorEnum struct {
	NO_TERMINATION ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior
	TERMINATE      ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior
}

func GetShowAutoLaunchGroupRespExcessFulfilledCapacityBehaviorEnum() ShowAutoLaunchGroupRespExcessFulfilledCapacityBehaviorEnum {
	return ShowAutoLaunchGroupRespExcessFulfilledCapacityBehaviorEnum{
		NO_TERMINATION: ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior{
			value: "noTermination",
		},
		TERMINATE: ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior{
			value: "terminate",
		},
	}
}

func (c ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior) Value() string {
	return c.value
}

func (c ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior) UnmarshalJSON(b []byte) error {
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

type ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration struct {
	value string
}

type ShowAutoLaunchGroupRespInstancesBehaviorWithExpirationEnum struct {
	TERMINATE      ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration
	NO_TERMINATION ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration
}

func GetShowAutoLaunchGroupRespInstancesBehaviorWithExpirationEnum() ShowAutoLaunchGroupRespInstancesBehaviorWithExpirationEnum {
	return ShowAutoLaunchGroupRespInstancesBehaviorWithExpirationEnum{
		TERMINATE: ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration{
			value: "terminate",
		},
		NO_TERMINATION: ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration{
			value: "noTermination",
		},
	}
}

func (c ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration) Value() string {
	return c.value
}

func (c ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration) UnmarshalJSON(b []byte) error {
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

type ShowAutoLaunchGroupRespAllocationStrategy struct {
	value string
}

type ShowAutoLaunchGroupRespAllocationStrategyEnum struct {
	LOWEST_PRICE       ShowAutoLaunchGroupRespAllocationStrategy
	DIVERSIFIED        ShowAutoLaunchGroupRespAllocationStrategy
	PRIORITIZED        ShowAutoLaunchGroupRespAllocationStrategy
	CAPACITY_OPTIMIZED ShowAutoLaunchGroupRespAllocationStrategy
}

func GetShowAutoLaunchGroupRespAllocationStrategyEnum() ShowAutoLaunchGroupRespAllocationStrategyEnum {
	return ShowAutoLaunchGroupRespAllocationStrategyEnum{
		LOWEST_PRICE: ShowAutoLaunchGroupRespAllocationStrategy{
			value: "lowest_price",
		},
		DIVERSIFIED: ShowAutoLaunchGroupRespAllocationStrategy{
			value: "diversified",
		},
		PRIORITIZED: ShowAutoLaunchGroupRespAllocationStrategy{
			value: "prioritized",
		},
		CAPACITY_OPTIMIZED: ShowAutoLaunchGroupRespAllocationStrategy{
			value: "capacity_optimized",
		},
	}
}

func (c ShowAutoLaunchGroupRespAllocationStrategy) Value() string {
	return c.value
}

func (c ShowAutoLaunchGroupRespAllocationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAutoLaunchGroupRespAllocationStrategy) UnmarshalJSON(b []byte) error {
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
