package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 智能购买组信息
type ShowAutoLaunchGroupResp struct {

	// autoLaunchGroup名称(1-64个字符)，只能包含中文、字母、数字、下划线或中划线
	Name string `json:"name"`

	// 请求类型 request：一次性。仅在启动时交付实例集群，调度失败后不再重试。 maintain：持续供应。在启动时尝试交付实例集群，并监控实时容量，未达到目标容量则尝试继续创建ECS实例
	Type ShowAutoLaunchGroupRespType `json:"type"`

	// AutoLaunchGroup状态,[SUBMITTED|ACTIVE|DELETING|DELETED]
	Status string `json:"status"`

	// 任务状态,[INIT|HANDLING|FULFILLED|ERROR]
	TaskState string `json:"task_state"`

	// 目标容量 实例数量或者CPU个数 目标容量大于等于stable_capacity， spot实例的容量为目标容量减去stable_capacity
	TargetCapacity int32 `json:"target_capacity"`

	// 按需实例目标容量 小于等于target_capacity
	StableCapacity int32 `json:"stable_capacity"`

	// 当前已经购买成功的总算力
	CurrentCapacity int32 `json:"current_capacity"`

	// 当前已经购买成功的按需容量
	CurrentStableCapacity int32 `json:"current_stable_capacity"`

	// 目标容量减少时 实例的中断行为 terminate|noTermination terminate：释放 noTermination：不释放 默认值：terminate
	ExcessFulfilledCapacityBehavior ShowAutoLaunchGroupRespExcessFulfilledCapacityBehavior `json:"excess_fulfilled_capacity_behavior"`

	// 请求到期 实例的中断行为 terminate|noTermination terminate：释放 noTermination：不释放 默认值：terminate
	InstancesBehaviorWithExpiration ShowAutoLaunchGroupRespInstancesBehaviorWithExpiration `json:"instances_behavior_with_expiration"`

	// 请求开始时间 yyyy-MM-dd HH:mm:ss
	ValidSince *sdktime.SdkTime `json:"valid_since"`

	// 请求结束时间 yyyy-MM-dd HH:mm:ss
	ValidUntil *sdktime.SdkTime `json:"valid_until"`

	// 智能购买组在各个region的配置
	RegionSpecs []RegionSpec `json:"region_specs"`

	// 实例分配策略 lowest_price:最小价格 diversified：AZ均铺 prioritized：优先级 capacity_optimized：容量最优化 默认值：lowest_price
	AllocationStrategy ShowAutoLaunchGroupRespAllocationStrategy `json:"allocation_strategy"`

	// spot实例价格 为全局spot实例的价格， 如果overrides中没有提供价格，可以使用该价格
	SpotPrice *float64 `json:"spot_price,omitempty"`

	// 智能购买组的创建时间
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
