package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFlavorSellPoliciesRequest struct {

	// 云服务器的系统规格的ID
	FlavorId *string `json:"flavor_id,omitempty"`

	// 云服务器的系统规格销售状态。  取值范围：  - available：正常售卖 - sellout：售罄
	SellStatus *ListFlavorSellPoliciesRequestSellStatus `json:"sell_status,omitempty"`

	// 计费模式。  key的取值范围：  - postPaid：按需计费实例。 - prePaid：包年/包月计费实例。 - spot：竞价实例。 - ri：预留实例。
	SellMode *ListFlavorSellPoliciesRequestSellMode `json:"sell_mode,omitempty"`

	// 可用区，需要指定可用区（AZ）
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 查询竞价实例时长大于设置值的策略
	LongestSpotDurationHoursGt *int32 `json:"longest_spot_duration_hours_gt,omitempty"`

	// 查询“竞价实例时长”的个数大于设置值的策略
	LargestSpotDurationCountGt *int32 `json:"largest_spot_duration_count_gt,omitempty"`

	// 查询竞价实例时长等于设置值的策略
	LongestSpotDurationHours *int32 `json:"longest_spot_duration_hours,omitempty"`

	// 查询“竞价实例时长”的个数等于设置值的策略
	LargestSpotDurationCount *int32 `json:"largest_spot_duration_count,omitempty"`

	// 中断策略。  取值范围：  - immediate：立即释放 - delay：延迟释放
	InterruptionPolicy *ListFlavorSellPoliciesRequestInterruptionPolicy `json:"interruption_policy,omitempty"`

	// 单页面可显示的flavor条数最大值，默认是1000。
	Limit *int32 `json:"limit,omitempty"`

	// 以单页最后一条flavor的ID作为分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListFlavorSellPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorSellPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorSellPoliciesRequest", string(data)}, " ")
}

type ListFlavorSellPoliciesRequestSellStatus struct {
	value string
}

type ListFlavorSellPoliciesRequestSellStatusEnum struct {
	AVAILABLE ListFlavorSellPoliciesRequestSellStatus
	SELLOUT   ListFlavorSellPoliciesRequestSellStatus
}

func GetListFlavorSellPoliciesRequestSellStatusEnum() ListFlavorSellPoliciesRequestSellStatusEnum {
	return ListFlavorSellPoliciesRequestSellStatusEnum{
		AVAILABLE: ListFlavorSellPoliciesRequestSellStatus{
			value: "available",
		},
		SELLOUT: ListFlavorSellPoliciesRequestSellStatus{
			value: "sellout",
		},
	}
}

func (c ListFlavorSellPoliciesRequestSellStatus) Value() string {
	return c.value
}

func (c ListFlavorSellPoliciesRequestSellStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorSellPoliciesRequestSellStatus) UnmarshalJSON(b []byte) error {
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

type ListFlavorSellPoliciesRequestSellMode struct {
	value string
}

type ListFlavorSellPoliciesRequestSellModeEnum struct {
	POST_PAID ListFlavorSellPoliciesRequestSellMode
	PRE_PAID  ListFlavorSellPoliciesRequestSellMode
	SPOT      ListFlavorSellPoliciesRequestSellMode
	RI        ListFlavorSellPoliciesRequestSellMode
}

func GetListFlavorSellPoliciesRequestSellModeEnum() ListFlavorSellPoliciesRequestSellModeEnum {
	return ListFlavorSellPoliciesRequestSellModeEnum{
		POST_PAID: ListFlavorSellPoliciesRequestSellMode{
			value: "postPaid",
		},
		PRE_PAID: ListFlavorSellPoliciesRequestSellMode{
			value: "prePaid",
		},
		SPOT: ListFlavorSellPoliciesRequestSellMode{
			value: "spot",
		},
		RI: ListFlavorSellPoliciesRequestSellMode{
			value: "ri",
		},
	}
}

func (c ListFlavorSellPoliciesRequestSellMode) Value() string {
	return c.value
}

func (c ListFlavorSellPoliciesRequestSellMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorSellPoliciesRequestSellMode) UnmarshalJSON(b []byte) error {
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

type ListFlavorSellPoliciesRequestInterruptionPolicy struct {
	value string
}

type ListFlavorSellPoliciesRequestInterruptionPolicyEnum struct {
	IMMEDIATE ListFlavorSellPoliciesRequestInterruptionPolicy
	DELAY     ListFlavorSellPoliciesRequestInterruptionPolicy
}

func GetListFlavorSellPoliciesRequestInterruptionPolicyEnum() ListFlavorSellPoliciesRequestInterruptionPolicyEnum {
	return ListFlavorSellPoliciesRequestInterruptionPolicyEnum{
		IMMEDIATE: ListFlavorSellPoliciesRequestInterruptionPolicy{
			value: "immediate",
		},
		DELAY: ListFlavorSellPoliciesRequestInterruptionPolicy{
			value: "delay",
		},
	}
}

func (c ListFlavorSellPoliciesRequestInterruptionPolicy) Value() string {
	return c.value
}

func (c ListFlavorSellPoliciesRequestInterruptionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorSellPoliciesRequestInterruptionPolicy) UnmarshalJSON(b []byte) error {
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
