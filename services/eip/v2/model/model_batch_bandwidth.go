package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchBandwidth 带宽信息
type BatchBandwidth struct {

	// - 功能说明：按流量计费还是按带宽计费。其中IPv6国外默认是bandwidth，国内默认是traffic。取值为traffic，表示流量计费。
	ChargeMode *BatchBandwidthChargeMode `json:"charge_mode,omitempty"`

	// - 功能说明：带宽名称 - 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点） - 约束：如果share_type是PER，该参数必须带,如果share_type是WHOLE并且id有值，该参数会忽略。
	Name *string `json:"name,omitempty"`

	// - 功能说明：带宽类型 - 取值范围：PER，WHOLE。其中IPv6暂不支持WHOLE类型带宽。
	ShareType *BatchBandwidthShareType `json:"share_type,omitempty"`

	// - 功能说明：带宽大小 - 取值范围：默认1Mbit/s~2000Mbit/s（具体范围以各区域配置为准，请参见控制台对应页面显示）。 - 约束：share_type是PER，该参数必须带，如果share_type是WHOLE并且id有值，该参数会忽略。 - 注意：调整带宽时的最小单位会根据带宽范围不同存在差异。   - 小于等于300Mbit/s：默认最小单位为1Mbit/s。   - 300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。   - 大于1000Mbit/s：默认最小单位为500Mbit/s。
	Size int32 `json:"size"`

	// - 功能说明：带宽ID，创建WHOLE类型带宽的弹性公网IP时可以指定之前的共享带宽创建 - 取值范围：WHOLE类型的带宽ID
	Id *string `json:"id,omitempty"`
}

func (o BatchBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBandwidth struct{}"
	}

	return strings.Join([]string{"BatchBandwidth", string(data)}, " ")
}

type BatchBandwidthChargeMode struct {
	value string
}

type BatchBandwidthChargeModeEnum struct {
	BANDWIDTH BatchBandwidthChargeMode
	TRAFFIC   BatchBandwidthChargeMode
}

func GetBatchBandwidthChargeModeEnum() BatchBandwidthChargeModeEnum {
	return BatchBandwidthChargeModeEnum{
		BANDWIDTH: BatchBandwidthChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: BatchBandwidthChargeMode{
			value: "traffic",
		},
	}
}

func (c BatchBandwidthChargeMode) Value() string {
	return c.value
}

func (c BatchBandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchBandwidthChargeMode) UnmarshalJSON(b []byte) error {
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

type BatchBandwidthShareType struct {
	value string
}

type BatchBandwidthShareTypeEnum struct {
	PER   BatchBandwidthShareType
	WHOLE BatchBandwidthShareType
}

func GetBatchBandwidthShareTypeEnum() BatchBandwidthShareTypeEnum {
	return BatchBandwidthShareTypeEnum{
		PER: BatchBandwidthShareType{
			value: "PER",
		},
		WHOLE: BatchBandwidthShareType{
			value: "WHOLE",
		},
	}
}

func (c BatchBandwidthShareType) Value() string {
	return c.value
}

func (c BatchBandwidthShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchBandwidthShareType) UnmarshalJSON(b []byte) error {
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
