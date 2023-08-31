package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateEquipmentResponse Response Object
type CreateEquipmentResponse struct {

	// 智能企业网关设备ID
	Id *string `json:"id,omitempty"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// esn
	Esn *string `json:"esn,omitempty"`

	// 设备名字
	Name *string `json:"name,omitempty"`

	// 设备类型
	Type *CreateEquipmentResponseType `json:"type,omitempty"`

	// 高可用类型
	HaType *CreateEquipmentResponseHaType `json:"ha_type,omitempty"`

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

func (o CreateEquipmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentResponse struct{}"
	}

	return strings.Join([]string{"CreateEquipmentResponse", string(data)}, " ")
}

type CreateEquipmentResponseType struct {
	value string
}

type CreateEquipmentResponseTypeEnum struct {
	STANDARD CreateEquipmentResponseType
}

func GetCreateEquipmentResponseTypeEnum() CreateEquipmentResponseTypeEnum {
	return CreateEquipmentResponseTypeEnum{
		STANDARD: CreateEquipmentResponseType{
			value: "standard",
		},
	}
}

func (c CreateEquipmentResponseType) Value() string {
	return c.value
}

func (c CreateEquipmentResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEquipmentResponseType) UnmarshalJSON(b []byte) error {
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

type CreateEquipmentResponseHaType struct {
	value string
}

type CreateEquipmentResponseHaTypeEnum struct {
	ACTIVE  CreateEquipmentResponseHaType
	STANDBY CreateEquipmentResponseHaType
}

func GetCreateEquipmentResponseHaTypeEnum() CreateEquipmentResponseHaTypeEnum {
	return CreateEquipmentResponseHaTypeEnum{
		ACTIVE: CreateEquipmentResponseHaType{
			value: "Active",
		},
		STANDBY: CreateEquipmentResponseHaType{
			value: "Standby",
		},
	}
}

func (c CreateEquipmentResponseHaType) Value() string {
	return c.value
}

func (c CreateEquipmentResponseHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEquipmentResponseHaType) UnmarshalJSON(b []byte) error {
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
