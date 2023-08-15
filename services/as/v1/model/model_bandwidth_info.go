package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BandwidthInfo 带宽信息
type BandwidthInfo struct {

	// 带宽（Mbit/s），按带宽分配取值范围为[1,2000]，按流量分配取值范围为[1,300]。  说明： - 若share_type是PER，该参数为必选项。若share_type是WHOLE，会忽略该参数。 - 具体范围以各区域配置为准，请参见控制台对应页面显示。 - 创建带宽时的最小单位会根据带宽取值范围不同存在差异。 - 小于等于300Mbit/s：默认最小单位为1Mbit/s。 - 300Mbit/s~1000Mbit/s：默认最小单位为50Mbit/s。 - 大于1000Mbit/s：默认最小单位为500Mbit/s。
	Size *int32 `json:"size,omitempty"`

	// 带宽的共享类型。共享类型枚举：PER：独享型。WHOLE：共享型。
	ShareType BandwidthInfoShareType `json:"share_type"`

	// 带宽的计费类型。字段值为“bandwidth”，表示按带宽计费。字段值为“traffic”，表示按流量计费。字段为其它值，会导致创建云服务器失败。如果share_type是PER，该参数为必选项。如果share_type是WHOLE，会忽略该参数。
	ChargingMode *BandwidthInfoChargingMode `json:"charging_mode,omitempty"`

	// 带宽ID，使用共享型带宽时，可以选择之前创建的共享带宽来创建弹性IP。如果share_type是PER，会忽略该参数。如果share_type是WHOLE，该参数为必选项。
	Id *string `json:"id,omitempty"`
}

func (o BandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfo struct{}"
	}

	return strings.Join([]string{"BandwidthInfo", string(data)}, " ")
}

type BandwidthInfoShareType struct {
	value string
}

type BandwidthInfoShareTypeEnum struct {
	PER   BandwidthInfoShareType
	WHOLE BandwidthInfoShareType
}

func GetBandwidthInfoShareTypeEnum() BandwidthInfoShareTypeEnum {
	return BandwidthInfoShareTypeEnum{
		PER: BandwidthInfoShareType{
			value: "PER",
		},
		WHOLE: BandwidthInfoShareType{
			value: "WHOLE",
		},
	}
}

func (c BandwidthInfoShareType) Value() string {
	return c.value
}

func (c BandwidthInfoShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthInfoShareType) UnmarshalJSON(b []byte) error {
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

type BandwidthInfoChargingMode struct {
	value string
}

type BandwidthInfoChargingModeEnum struct {
	BANDWIDTH BandwidthInfoChargingMode
	TRAFFIC   BandwidthInfoChargingMode
}

func GetBandwidthInfoChargingModeEnum() BandwidthInfoChargingModeEnum {
	return BandwidthInfoChargingModeEnum{
		BANDWIDTH: BandwidthInfoChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: BandwidthInfoChargingMode{
			value: "traffic",
		},
	}
}

func (c BandwidthInfoChargingMode) Value() string {
	return c.value
}

func (c BandwidthInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthInfoChargingMode) UnmarshalJSON(b []byte) error {
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
