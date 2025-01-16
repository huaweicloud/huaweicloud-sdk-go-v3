package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceBandwidthAutoScalingPolicyResponse Response Object
type ShowInstanceBandwidthAutoScalingPolicyResponse struct {

	// 带宽弹性的观测窗口，单位：分钟。支持的取值：1、5、10、15、30。
	WindowSize ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize `json:"window_size"`

	// 触发带宽自动扩展的带宽平均使用率阈值，单位：百分比。支持的取值：50、60、70、80、90、95。
	BandwidthUsageUpperThreshold ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold `json:"bandwidth_usage_upper_threshold"`

	// 带宽扩展操作的静默时间（两次带宽扩展操作之间的最小间隔时间），单位：秒。 默认值：0。取值范围：0~86400。
	ScaleOutCooldown int32 `json:"scale_out_cooldown"`

	// 是否启用带宽自动回缩。默认值：false。该参数暂未启用。
	ScaleInEnabled *bool `json:"scale_in_enabled,omitempty"`

	// 触发带宽自动回缩的带宽平均使用率阈值，单位：百分比。支持的取值：10、20、30。该参数暂未启用。
	BandwidthUsageLowerThreshold *int32 `json:"bandwidth_usage_lower_threshold,omitempty"`

	// 带宽回缩操作的静默时间（两次带宽回缩操作之间的最小间隔时间），单位：秒。该参数暂未启用。 默认值：300。取值范围：0~86400。
	ScaleInCooldown *int32 `json:"scale_in_cooldown,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowInstanceBandwidthAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBandwidthAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceBandwidthAutoScalingPolicyResponse", string(data)}, " ")
}

type ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize struct {
	value int32
}

type ShowInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum struct {
	E_1  ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_5  ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_10 ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_15 ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_30 ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize
}

func GetShowInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum() ShowInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum {
	return ShowInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum{
		E_1: ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 1,
		}, E_5: ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 5,
		}, E_10: ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 10,
		}, E_15: ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 15,
		}, E_30: ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 30,
		},
	}
}

func (c ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize) Value() int32 {
	return c.value
}

func (c ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceBandwidthAutoScalingPolicyResponseWindowSize) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold struct {
	value int32
}

type ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum struct {
	E_50 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_60 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_70 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_80 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_90 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_95 ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
}

func GetShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum() ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum {
	return ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum{
		E_50: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 50,
		}, E_60: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 60,
		}, E_70: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 70,
		}, E_80: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 80,
		}, E_90: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 90,
		}, E_95: ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 95,
		},
	}
}

func (c ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) Value() int32 {
	return c.value
}

func (c ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
