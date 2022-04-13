package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 带宽对象
type Bandwidth struct {
	// 带宽类型。 取值范围： share：共享类型

	BandwidthType *BandwidthBandwidthType `json:"bandwidth_type,omitempty"`
	// 计费模式，当前只支持峰值95计费。  取值范围：  - 95peak_plus：峰值95计费

	ChargeMode *BandwidthChargeMode `json:"charge_mode,omitempty"`
	// 创建时间。

	CreateTime *string `json:"create_time,omitempty"`
	// 带宽ID。

	Id *string `json:"id,omitempty"`
	// 带宽名称。

	Name *string `json:"name,omitempty"`

	Operator *Operator `json:"operator,omitempty"`
	// 弹性公网IP信息。

	PublicipInfo *[]PublicipInfo `json:"publicip_info,omitempty"`
	// 共享带宽类型，标识是否是共享带宽。  取值范围：  - WHOLE：共享带宽

	ShareType *BandwidthShareType `json:"share_type,omitempty"`
	// 边缘站点ID。

	SiteId *string `json:"site_id,omitempty"`
	// 站点信息。

	SiteInfo *string `json:"site_info,omitempty"`
	// 带宽大小。

	Size *int32 `json:"size,omitempty"`
	// 带宽的状态。  取值范围：  - FREEZED：冻结  - NORMAL：正常

	Status *BandwidthStatus `json:"status,omitempty"`
	// 更新时间。

	UpdateTime *string `json:"update_time,omitempty"`
	// 线路ID。

	PoolId *string `json:"pool_id,omitempty"`
}

func (o Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Bandwidth struct{}"
	}

	return strings.Join([]string{"Bandwidth", string(data)}, " ")
}

type BandwidthBandwidthType struct {
	value string
}

type BandwidthBandwidthTypeEnum struct {
	SHARE BandwidthBandwidthType
}

func GetBandwidthBandwidthTypeEnum() BandwidthBandwidthTypeEnum {
	return BandwidthBandwidthTypeEnum{
		SHARE: BandwidthBandwidthType{
			value: "share",
		},
	}
}

func (c BandwidthBandwidthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthBandwidthType) UnmarshalJSON(b []byte) error {
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

type BandwidthChargeMode struct {
	value string
}

type BandwidthChargeModeEnum struct {
	E_95PEAK_PLUS BandwidthChargeMode
}

func GetBandwidthChargeModeEnum() BandwidthChargeModeEnum {
	return BandwidthChargeModeEnum{
		E_95PEAK_PLUS: BandwidthChargeMode{
			value: "95peak_plus",
		},
	}
}

func (c BandwidthChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthChargeMode) UnmarshalJSON(b []byte) error {
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

type BandwidthShareType struct {
	value string
}

type BandwidthShareTypeEnum struct {
	WHOLE BandwidthShareType
}

func GetBandwidthShareTypeEnum() BandwidthShareTypeEnum {
	return BandwidthShareTypeEnum{
		WHOLE: BandwidthShareType{
			value: "WHOLE",
		},
	}
}

func (c BandwidthShareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthShareType) UnmarshalJSON(b []byte) error {
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

type BandwidthStatus struct {
	value string
}

type BandwidthStatusEnum struct {
	FREEZED BandwidthStatus
	NORMAL  BandwidthStatus
}

func GetBandwidthStatusEnum() BandwidthStatusEnum {
	return BandwidthStatusEnum{
		FREEZED: BandwidthStatus{
			value: "FREEZED",
		},
		NORMAL: BandwidthStatus{
			value: "NORMAL",
		},
	}
}

func (c BandwidthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthStatus) UnmarshalJSON(b []byte) error {
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
