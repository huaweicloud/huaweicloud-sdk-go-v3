package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VifExtendAttribute 接口BFD/NQA等可靠性检测信息
type VifExtendAttribute struct {

	// 虚拟接口的可用性检测类型
	HaType *VifExtendAttributeHaType `json:"ha_type,omitempty"`

	// 检测的具体的配置模式
	HaMode *VifExtendAttributeHaMode `json:"ha_mode,omitempty"`

	// 检测的重试次数
	DetectMultiplier *int32 `json:"detect_multiplier,omitempty"`

	// 检测的接收时长间隔
	MinRxInterval *int32 `json:"min_rx_interval,omitempty"`

	// 检测的发送时长间隔
	MinTxInterval *int32 `json:"min_tx_interval,omitempty"`

	// 检测的远端的标识，用于静态BFD
	RemoteDisclaim *int32 `json:"remote_disclaim,omitempty"`

	// 检测的本端的标识，用于静态BFD
	LocalDisclaim *int32 `json:"local_disclaim,omitempty"`
}

func (o VifExtendAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VifExtendAttribute struct{}"
	}

	return strings.Join([]string{"VifExtendAttribute", string(data)}, " ")
}

type VifExtendAttributeHaType struct {
	value string
}

type VifExtendAttributeHaTypeEnum struct {
	NQA VifExtendAttributeHaType
	BFD VifExtendAttributeHaType
}

func GetVifExtendAttributeHaTypeEnum() VifExtendAttributeHaTypeEnum {
	return VifExtendAttributeHaTypeEnum{
		NQA: VifExtendAttributeHaType{
			value: "nqa",
		},
		BFD: VifExtendAttributeHaType{
			value: "bfd",
		},
	}
}

func (c VifExtendAttributeHaType) Value() string {
	return c.value
}

func (c VifExtendAttributeHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VifExtendAttributeHaType) UnmarshalJSON(b []byte) error {
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

type VifExtendAttributeHaMode struct {
	value string
}

type VifExtendAttributeHaModeEnum struct {
	AUTO_SINGLE   VifExtendAttributeHaMode
	AUTO_MULTI    VifExtendAttributeHaMode
	STATIC_SINGLE VifExtendAttributeHaMode
	STATIC_MULTI  VifExtendAttributeHaMode
	ENHANCE_NQA   VifExtendAttributeHaMode
}

func GetVifExtendAttributeHaModeEnum() VifExtendAttributeHaModeEnum {
	return VifExtendAttributeHaModeEnum{
		AUTO_SINGLE: VifExtendAttributeHaMode{
			value: "auto_single",
		},
		AUTO_MULTI: VifExtendAttributeHaMode{
			value: "auto_multi",
		},
		STATIC_SINGLE: VifExtendAttributeHaMode{
			value: "static_single",
		},
		STATIC_MULTI: VifExtendAttributeHaMode{
			value: "static_multi",
		},
		ENHANCE_NQA: VifExtendAttributeHaMode{
			value: "enhance_nqa",
		},
	}
}

func (c VifExtendAttributeHaMode) Value() string {
	return c.value
}

func (c VifExtendAttributeHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VifExtendAttributeHaMode) UnmarshalJSON(b []byte) error {
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
