package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwitchWebTamperProtectStatusRequestBody struct {

	// 防护类型
	Type SwitchWebTamperProtectStatusRequestBodyType `json:"type"`

	// 操作类型：（open 开启防护,close 关闭防护）
	OperateType SwitchWebTamperProtectStatusRequestBodyOperateType `json:"operate_type"`

	// 计费模式，开启防护时需要，容器网页防篡改固定传packet_cycle
	ChargingMode SwitchWebTamperProtectStatusRequestBodyChargingMode `json:"charging_mode"`

	// 操作的防护配置id列表
	ProtectionConfigIdList []string `json:"protection_config_id_list"`

	// 配额id列表，开启防护时需要（若开启防护时不传该列表，则随机绑定配额）
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`
}

func (o SwitchWebTamperProtectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWebTamperProtectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchWebTamperProtectStatusRequestBody", string(data)}, " ")
}

type SwitchWebTamperProtectStatusRequestBodyType struct {
	value string
}

type SwitchWebTamperProtectStatusRequestBodyTypeEnum struct {
	CONTAINER_WTP SwitchWebTamperProtectStatusRequestBodyType
}

func GetSwitchWebTamperProtectStatusRequestBodyTypeEnum() SwitchWebTamperProtectStatusRequestBodyTypeEnum {
	return SwitchWebTamperProtectStatusRequestBodyTypeEnum{
		CONTAINER_WTP: SwitchWebTamperProtectStatusRequestBodyType{
			value: "container_wtp",
		},
	}
}

func (c SwitchWebTamperProtectStatusRequestBodyType) Value() string {
	return c.value
}

func (c SwitchWebTamperProtectStatusRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchWebTamperProtectStatusRequestBodyType) UnmarshalJSON(b []byte) error {
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

type SwitchWebTamperProtectStatusRequestBodyOperateType struct {
	value string
}

type SwitchWebTamperProtectStatusRequestBodyOperateTypeEnum struct {
	CLOSE SwitchWebTamperProtectStatusRequestBodyOperateType
	OPEN  SwitchWebTamperProtectStatusRequestBodyOperateType
}

func GetSwitchWebTamperProtectStatusRequestBodyOperateTypeEnum() SwitchWebTamperProtectStatusRequestBodyOperateTypeEnum {
	return SwitchWebTamperProtectStatusRequestBodyOperateTypeEnum{
		CLOSE: SwitchWebTamperProtectStatusRequestBodyOperateType{
			value: "close",
		},
		OPEN: SwitchWebTamperProtectStatusRequestBodyOperateType{
			value: "open",
		},
	}
}

func (c SwitchWebTamperProtectStatusRequestBodyOperateType) Value() string {
	return c.value
}

func (c SwitchWebTamperProtectStatusRequestBodyOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchWebTamperProtectStatusRequestBodyOperateType) UnmarshalJSON(b []byte) error {
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

type SwitchWebTamperProtectStatusRequestBodyChargingMode struct {
	value string
}

type SwitchWebTamperProtectStatusRequestBodyChargingModeEnum struct {
	PACKET_CYCLE SwitchWebTamperProtectStatusRequestBodyChargingMode
}

func GetSwitchWebTamperProtectStatusRequestBodyChargingModeEnum() SwitchWebTamperProtectStatusRequestBodyChargingModeEnum {
	return SwitchWebTamperProtectStatusRequestBodyChargingModeEnum{
		PACKET_CYCLE: SwitchWebTamperProtectStatusRequestBodyChargingMode{
			value: "packet_cycle",
		},
	}
}

func (c SwitchWebTamperProtectStatusRequestBodyChargingMode) Value() string {
	return c.value
}

func (c SwitchWebTamperProtectStatusRequestBodyChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchWebTamperProtectStatusRequestBodyChargingMode) UnmarshalJSON(b []byte) error {
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
