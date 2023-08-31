package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateEquipmentInfoResponse Response Object
type UpdateEquipmentInfoResponse struct {

	// 智能企业网关设备ID
	Id *string `json:"id,omitempty"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// esn
	Esn *string `json:"esn,omitempty"`

	// 设备名字
	Name *string `json:"name,omitempty"`

	// 设备类型
	Type *UpdateEquipmentInfoResponseType `json:"type,omitempty"`

	// 高可用类型
	HaType *UpdateEquipmentInfoResponseHaType `json:"ha_type,omitempty"`

	// 设备软件版本
	Version *string `json:"version,omitempty"`

	// 激活时间
	ActiveAt *sdktime.SdkTime `json:"active_at,omitempty"`

	// 上线时间
	GoLiveAt *sdktime.SdkTime `json:"go_live_at,omitempty"`

	// 设备启动时间
	StartUpAt *sdktime.SdkTime `json:"start_up_at,omitempty"`

	// VPN状态
	CloudAccessStatus *string `json:"cloud_access_status,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEquipmentInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentInfoResponse", string(data)}, " ")
}

type UpdateEquipmentInfoResponseType struct {
	value string
}

type UpdateEquipmentInfoResponseTypeEnum struct {
	STANDARD UpdateEquipmentInfoResponseType
}

func GetUpdateEquipmentInfoResponseTypeEnum() UpdateEquipmentInfoResponseTypeEnum {
	return UpdateEquipmentInfoResponseTypeEnum{
		STANDARD: UpdateEquipmentInfoResponseType{
			value: "standard",
		},
	}
}

func (c UpdateEquipmentInfoResponseType) Value() string {
	return c.value
}

func (c UpdateEquipmentInfoResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEquipmentInfoResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateEquipmentInfoResponseHaType struct {
	value string
}

type UpdateEquipmentInfoResponseHaTypeEnum struct {
	ACTIVE  UpdateEquipmentInfoResponseHaType
	STANDBY UpdateEquipmentInfoResponseHaType
}

func GetUpdateEquipmentInfoResponseHaTypeEnum() UpdateEquipmentInfoResponseHaTypeEnum {
	return UpdateEquipmentInfoResponseHaTypeEnum{
		ACTIVE: UpdateEquipmentInfoResponseHaType{
			value: "Active",
		},
		STANDBY: UpdateEquipmentInfoResponseHaType{
			value: "Standby",
		},
	}
}

func (c UpdateEquipmentInfoResponseHaType) Value() string {
	return c.value
}

func (c UpdateEquipmentInfoResponseHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEquipmentInfoResponseHaType) UnmarshalJSON(b []byte) error {
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
