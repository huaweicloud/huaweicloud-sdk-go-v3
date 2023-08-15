package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubMetricsRequest Request Object
type ListSubMetricsRequest struct {

	// 指标数据统计方式
	Filter *ListSubMetricsRequestFilter `json:"filter,omitempty"`

	// 指标数据统计周期（单位minute）
	Period *int32 `json:"period,omitempty"`

	// 获取指标数据起始时间
	StartTime int64 `json:"start_time"`

	// 获取指标数据结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 事件订阅id
	SubscriptionId string `json:"subscription_id"`

	// 事件目标类型/事件通道类型
	ProviderType *ListSubMetricsRequestProviderType `json:"provider_type,omitempty"`

	// 事件目标id
	TargetId *string `json:"target_id,omitempty"`
}

func (o ListSubMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListSubMetricsRequest", string(data)}, " ")
}

type ListSubMetricsRequestFilter struct {
	value string
}

type ListSubMetricsRequestFilterEnum struct {
	AVG ListSubMetricsRequestFilter
	MIN ListSubMetricsRequestFilter
	MAX ListSubMetricsRequestFilter
}

func GetListSubMetricsRequestFilterEnum() ListSubMetricsRequestFilterEnum {
	return ListSubMetricsRequestFilterEnum{
		AVG: ListSubMetricsRequestFilter{
			value: "AVG",
		},
		MIN: ListSubMetricsRequestFilter{
			value: "MIN",
		},
		MAX: ListSubMetricsRequestFilter{
			value: "MAX",
		},
	}
}

func (c ListSubMetricsRequestFilter) Value() string {
	return c.value
}

func (c ListSubMetricsRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubMetricsRequestFilter) UnmarshalJSON(b []byte) error {
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

type ListSubMetricsRequestProviderType struct {
	value string
}

type ListSubMetricsRequestProviderTypeEnum struct {
	OFFICIAL ListSubMetricsRequestProviderType
	CUSTOM   ListSubMetricsRequestProviderType
}

func GetListSubMetricsRequestProviderTypeEnum() ListSubMetricsRequestProviderTypeEnum {
	return ListSubMetricsRequestProviderTypeEnum{
		OFFICIAL: ListSubMetricsRequestProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ListSubMetricsRequestProviderType{
			value: "CUSTOM",
		},
	}
}

func (c ListSubMetricsRequestProviderType) Value() string {
	return c.value
}

func (c ListSubMetricsRequestProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubMetricsRequestProviderType) UnmarshalJSON(b []byte) error {
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
