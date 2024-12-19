package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateInstanceBandwidthAutoScalingPolicyRequestBody 更新实例带宽弹性伸缩策略请求体。
type UpdateInstanceBandwidthAutoScalingPolicyRequestBody struct {

	// 带宽弹性的观测窗口，单位：分钟。支持的取值：1、5、10、15、30。
	WindowSize UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize `json:"window_size"`

	// 触发带宽自动扩展的带宽平均使用率阈值，单位：百分比。支持的取值：50、60、70、80、90、95。
	BandwidthUsageUpperThreshold UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold `json:"bandwidth_usage_upper_threshold"`

	// 带宽扩展操作的静默时间（两次带宽扩展操作之间的最小间隔时间），单位：秒。 默认值：0。取值范围：0~86400。
	ScaleOutCooldown *int32 `json:"scale_out_cooldown,omitempty"`

	// 是否启用带宽自动回缩。默认值：false。该参数暂未启用。
	ScaleInEnabled *bool `json:"scale_in_enabled,omitempty"`

	// 触发带宽自动回缩的带宽平均使用率阈值，单位：百分比。支持的取值：10、20、30。该参数暂未启用。
	BandwidthUsageLowerThreshold *int32 `json:"bandwidth_usage_lower_threshold,omitempty"`

	// 带宽回缩操作的静默时间（两次带宽回缩操作之间的最小间隔时间），单位：秒。该参数暂未启用。 默认值：300。取值范围：0~86400。
	ScaleInCooldown *int32 `json:"scale_in_cooldown,omitempty"`
}

func (o UpdateInstanceBandwidthAutoScalingPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBandwidthAutoScalingPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBandwidthAutoScalingPolicyRequestBody", string(data)}, " ")
}

type UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize struct {
	value int32
}

type UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSizeEnum struct {
	E_1  UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize
	E_5  UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize
	E_10 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize
	E_15 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize
	E_30 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize
}

func GetUpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSizeEnum() UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSizeEnum {
	return UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSizeEnum{
		E_1: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize{
			value: 1,
		}, E_5: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize{
			value: 5,
		}, E_10: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize{
			value: 10,
		}, E_15: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize{
			value: 15,
		}, E_30: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize{
			value: 30,
		},
	}
}

func (c UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize) Value() int32 {
	return c.value
}

func (c UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceBandwidthAutoScalingPolicyRequestBodyWindowSize) UnmarshalJSON(b []byte) error {
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

type UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold struct {
	value int32
}

type UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThresholdEnum struct {
	E_50 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
	E_60 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
	E_70 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
	E_80 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
	E_90 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
	E_95 UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold
}

func GetUpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThresholdEnum() UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThresholdEnum {
	return UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThresholdEnum{
		E_50: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 50,
		}, E_60: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 60,
		}, E_70: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 70,
		}, E_80: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 80,
		}, E_90: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 90,
		}, E_95: UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold{
			value: 95,
		},
	}
}

func (c UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold) Value() int32 {
	return c.value
}

func (c UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceBandwidthAutoScalingPolicyRequestBodyBandwidthUsageUpperThreshold) UnmarshalJSON(b []byte) error {
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
