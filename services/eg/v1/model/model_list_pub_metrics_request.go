package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListPubMetricsRequest struct {

	// 指标数据统计方式
	Filter *ListPubMetricsRequestFilter `json:"filter,omitempty"`

	// 指标数据统计周期（单位minute）
	Period *int32 `json:"period,omitempty"`

	// 获取指标数据起始时间
	StartTime int64 `json:"start_time"`

	// 获取指标数据结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 事件通道id
	ChannelId string `json:"channel_id"`

	// 事件目标类型/事件通道类型
	ProviderType *ListPubMetricsRequestProviderType `json:"provider_type,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`
}

func (o ListPubMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPubMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListPubMetricsRequest", string(data)}, " ")
}

type ListPubMetricsRequestFilter struct {
	value string
}

type ListPubMetricsRequestFilterEnum struct {
	AVG ListPubMetricsRequestFilter
	MIN ListPubMetricsRequestFilter
	MAX ListPubMetricsRequestFilter
}

func GetListPubMetricsRequestFilterEnum() ListPubMetricsRequestFilterEnum {
	return ListPubMetricsRequestFilterEnum{
		AVG: ListPubMetricsRequestFilter{
			value: "AVG",
		},
		MIN: ListPubMetricsRequestFilter{
			value: "MIN",
		},
		MAX: ListPubMetricsRequestFilter{
			value: "MAX",
		},
	}
}

func (c ListPubMetricsRequestFilter) Value() string {
	return c.value
}

func (c ListPubMetricsRequestFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPubMetricsRequestFilter) UnmarshalJSON(b []byte) error {
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

type ListPubMetricsRequestProviderType struct {
	value string
}

type ListPubMetricsRequestProviderTypeEnum struct {
	OFFICIAL ListPubMetricsRequestProviderType
	CUSTOM   ListPubMetricsRequestProviderType
}

func GetListPubMetricsRequestProviderTypeEnum() ListPubMetricsRequestProviderTypeEnum {
	return ListPubMetricsRequestProviderTypeEnum{
		OFFICIAL: ListPubMetricsRequestProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ListPubMetricsRequestProviderType{
			value: "CUSTOM",
		},
	}
}

func (c ListPubMetricsRequestProviderType) Value() string {
	return c.value
}

func (c ListPubMetricsRequestProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPubMetricsRequestProviderType) UnmarshalJSON(b []byte) error {
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
