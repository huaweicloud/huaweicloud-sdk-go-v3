package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthAutoScalingPolicy 带宽弹性伸缩策略。
type BandwidthAutoScalingPolicy struct {

	// 带宽弹性的观测窗口，单位：分钟。支持的取值：1、5、10、15、30。
	WindowSize BandwidthAutoScalingPolicyWindowSize `json:"window_size"`

	// 触发带宽自动扩展的带宽平均使用率阈值，单位：百分比。支持的取值：50、60、70、80、90、95。
	BandwidthUsageUpperThreshold BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold `json:"bandwidth_usage_upper_threshold"`

	// 带宽扩展操作的静默时间（两次带宽扩展操作之间的最小间隔时间），单位：秒。 默认值：0。取值范围：0~86400。
	ScaleOutCooldown int32 `json:"scale_out_cooldown"`

	// 是否启用带宽自动回缩。默认值：false。该参数暂未启用。
	ScaleInEnabled *bool `json:"scale_in_enabled,omitempty"`

	// 触发带宽自动回缩的带宽平均使用率阈值，单位：百分比。支持的取值：10、20、30。该参数暂未启用。
	BandwidthUsageLowerThreshold *int32 `json:"bandwidth_usage_lower_threshold,omitempty"`

	// 带宽回缩操作的静默时间（两次带宽回缩操作之间的最小间隔时间），单位：秒。该参数暂未启用。 默认值：300。取值范围：0~86400。
	ScaleInCooldown *int32 `json:"scale_in_cooldown,omitempty"`
}

func (o BandwidthAutoScalingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthAutoScalingPolicy struct{}"
	}

	return strings.Join([]string{"BandwidthAutoScalingPolicy", string(data)}, " ")
}

type BandwidthAutoScalingPolicyWindowSize struct {
	value int32
}

type BandwidthAutoScalingPolicyWindowSizeEnum struct {
	E_1  BandwidthAutoScalingPolicyWindowSize
	E_5  BandwidthAutoScalingPolicyWindowSize
	E_10 BandwidthAutoScalingPolicyWindowSize
	E_15 BandwidthAutoScalingPolicyWindowSize
	E_30 BandwidthAutoScalingPolicyWindowSize
}

func GetBandwidthAutoScalingPolicyWindowSizeEnum() BandwidthAutoScalingPolicyWindowSizeEnum {
	return BandwidthAutoScalingPolicyWindowSizeEnum{
		E_1: BandwidthAutoScalingPolicyWindowSize{
			value: 1,
		}, E_5: BandwidthAutoScalingPolicyWindowSize{
			value: 5,
		}, E_10: BandwidthAutoScalingPolicyWindowSize{
			value: 10,
		}, E_15: BandwidthAutoScalingPolicyWindowSize{
			value: 15,
		}, E_30: BandwidthAutoScalingPolicyWindowSize{
			value: 30,
		},
	}
}

func (c BandwidthAutoScalingPolicyWindowSize) Value() int32 {
	return c.value
}

func (c BandwidthAutoScalingPolicyWindowSize) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthAutoScalingPolicyWindowSize) UnmarshalJSON(b []byte) error {
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

type BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold struct {
	value int32
}

type BandwidthAutoScalingPolicyBandwidthUsageUpperThresholdEnum struct {
	E_50 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
	E_60 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
	E_70 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
	E_80 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
	E_90 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
	E_95 BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold
}

func GetBandwidthAutoScalingPolicyBandwidthUsageUpperThresholdEnum() BandwidthAutoScalingPolicyBandwidthUsageUpperThresholdEnum {
	return BandwidthAutoScalingPolicyBandwidthUsageUpperThresholdEnum{
		E_50: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 50,
		}, E_60: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 60,
		}, E_70: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 70,
		}, E_80: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 80,
		}, E_90: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 90,
		}, E_95: BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold{
			value: 95,
		},
	}
}

func (c BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold) Value() int32 {
	return c.value
}

func (c BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthAutoScalingPolicyBandwidthUsageUpperThreshold) UnmarshalJSON(b []byte) error {
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
