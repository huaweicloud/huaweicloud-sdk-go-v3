package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 带宽信息
type CreateLoadBalancerBandwidthOption struct {

	// 带宽名称。  取值：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）  使用说明： - 如果share_type是PER，该字段是必选。 - 如果bandwidth对象的id有值，该字段被忽略。
	Name *string `json:"name,omitempty"`

	// 带宽大小 取值范围:默认1Mbit/s~2000Mbit/s(具体范围以各区域配置为准,请参见控制台对应页面显示)。  注意:调整带宽时的最小单位会根据带宽范围不同存在差异。 小于等于300Mbit/s:默认最小单位为1Mbit/s。 300Mbit/s~1000Mbit/s:默认最小单位为50Mbit/s。 大于1000Mbit/s:默认最小单位为500Mbit/s。  使用说明：当id字段为null时，size是必须的。
	Size *int32 `json:"size,omitempty"`

	// 计费模式。  [取值：bandwidth表示按带宽计费，traffic表示按流量计费。 ](tag:hws,hws_hk,ocb,ctc,hcs,tm,cmcc,hws_ocb,fcs)  [当前仅支持traffic按流量计费](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt)  使用说明：当id字段为null时，charge_mode是必须的。
	ChargeMode *CreateLoadBalancerBandwidthOptionChargeMode `json:"charge_mode,omitempty"`

	// 带宽类型。  取值： - PER：独享带宽。 - WHOLE：共享带宽。  使用说明： - 当id字段为null时，share_type是必须的。当id不为null时，该字段被忽略。 - 该字段为WHOLE时,必须指定带宽ID。 - IPv6的EIP不支持WHOLE类型带宽。
	ShareType *CreateLoadBalancerBandwidthOptionShareType `json:"share_type,omitempty"`

	// 资源账单信息。  [如果billing_info不为空，说明是包周期计费的带宽，否则为按需计费的带宽 ](tag:hws,hws_hk,ocb,tlf,ctc,hcso,sbc,tm,cmcc)  [不支持该字段，请勿使用](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt)
	BillingInfo *string `json:"billing_info,omitempty"`

	// 功能说明：使用已有的共享带宽创建IP 取值范围：共享带宽ID 使用说明： WHOLE类型的带宽ID； 在预付费的情况下，不填该值。该字段取空字符串时，会被忽略。
	Id *string `json:"id,omitempty"`
}

func (o CreateLoadBalancerBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerBandwidthOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerBandwidthOption", string(data)}, " ")
}

type CreateLoadBalancerBandwidthOptionChargeMode struct {
	value string
}

type CreateLoadBalancerBandwidthOptionChargeModeEnum struct {
	BANDWIDTH CreateLoadBalancerBandwidthOptionChargeMode
	TRAFFIC   CreateLoadBalancerBandwidthOptionChargeMode
}

func GetCreateLoadBalancerBandwidthOptionChargeModeEnum() CreateLoadBalancerBandwidthOptionChargeModeEnum {
	return CreateLoadBalancerBandwidthOptionChargeModeEnum{
		BANDWIDTH: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: CreateLoadBalancerBandwidthOptionChargeMode{
			value: "traffic",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionChargeMode) Value() string {
	return c.value
}

func (c CreateLoadBalancerBandwidthOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateLoadBalancerBandwidthOptionShareType struct {
	value string
}

type CreateLoadBalancerBandwidthOptionShareTypeEnum struct {
	PER   CreateLoadBalancerBandwidthOptionShareType
	WHOLE CreateLoadBalancerBandwidthOptionShareType
}

func GetCreateLoadBalancerBandwidthOptionShareTypeEnum() CreateLoadBalancerBandwidthOptionShareTypeEnum {
	return CreateLoadBalancerBandwidthOptionShareTypeEnum{
		PER: CreateLoadBalancerBandwidthOptionShareType{
			value: "PER",
		},
		WHOLE: CreateLoadBalancerBandwidthOptionShareType{
			value: "WHOLE",
		},
	}
}

func (c CreateLoadBalancerBandwidthOptionShareType) Value() string {
	return c.value
}

func (c CreateLoadBalancerBandwidthOptionShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerBandwidthOptionShareType) UnmarshalJSON(b []byte) error {
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
