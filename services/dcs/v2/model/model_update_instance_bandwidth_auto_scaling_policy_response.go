package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateInstanceBandwidthAutoScalingPolicyResponse Response Object
type UpdateInstanceBandwidthAutoScalingPolicyResponse struct {

	// 带宽弹性的观测窗口，单位：分钟。支持的取值：1、5、10、15、30。
	WindowSize UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize `json:"window_size"`

	// 触发带宽自动扩展的带宽平均使用率阈值，单位：百分比。支持的取值：50、60、70、80、90、95。
	BandwidthUsageUpperThreshold UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold `json:"bandwidth_usage_upper_threshold"`

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

func (o UpdateInstanceBandwidthAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBandwidthAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBandwidthAutoScalingPolicyResponse", string(data)}, " ")
}

type UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize struct {
	value int32
}

type UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum struct {
	E_1  UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_5  UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_10 UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_15 UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize
	E_30 UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize
}

func GetUpdateInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum() UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum {
	return UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSizeEnum{
		E_1: UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 1,
		}, E_5: UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 5,
		}, E_10: UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 10,
		}, E_15: UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 15,
		}, E_30: UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize{
			value: 30,
		},
	}
}

func (c UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize) Value() int32 {
	return c.value
}

func (c UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceBandwidthAutoScalingPolicyResponseWindowSize) UnmarshalJSON(b []byte) error {
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

type UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold struct {
	value int32
}

type UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum struct {
	E_50 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_60 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_70 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_80 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_90 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
	E_95 UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold
}

func GetUpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum() UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum {
	return UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThresholdEnum{
		E_50: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 50,
		}, E_60: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 60,
		}, E_70: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 70,
		}, E_80: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 80,
		}, E_90: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 90,
		}, E_95: UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold{
			value: 95,
		},
	}
}

func (c UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) Value() int32 {
	return c.value
}

func (c UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceBandwidthAutoScalingPolicyResponseBandwidthUsageUpperThreshold) UnmarshalJSON(b []byte) error {
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
