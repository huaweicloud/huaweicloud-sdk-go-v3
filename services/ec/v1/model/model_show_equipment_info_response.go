package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowEquipmentInfoResponse Response Object
type ShowEquipmentInfoResponse struct {

	// 智能企业网关设备ID
	Id *string `json:"id,omitempty"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// esn
	Esn *string `json:"esn,omitempty"`

	// 设备名字
	Name *string `json:"name,omitempty"`

	// 设备类型
	Type *ShowEquipmentInfoResponseType `json:"type,omitempty"`

	// 高可用类型
	HaType *ShowEquipmentInfoResponseHaType `json:"ha_type,omitempty"`

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

func (o ShowEquipmentInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEquipmentInfoResponse", string(data)}, " ")
}

type ShowEquipmentInfoResponseType struct {
	value string
}

type ShowEquipmentInfoResponseTypeEnum struct {
	STANDARD ShowEquipmentInfoResponseType
	SOHO     ShowEquipmentInfoResponseType
}

func GetShowEquipmentInfoResponseTypeEnum() ShowEquipmentInfoResponseTypeEnum {
	return ShowEquipmentInfoResponseTypeEnum{
		STANDARD: ShowEquipmentInfoResponseType{
			value: "standard",
		},
		SOHO: ShowEquipmentInfoResponseType{
			value: "soho",
		},
	}
}

func (c ShowEquipmentInfoResponseType) Value() string {
	return c.value
}

func (c ShowEquipmentInfoResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEquipmentInfoResponseType) UnmarshalJSON(b []byte) error {
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

type ShowEquipmentInfoResponseHaType struct {
	value string
}

type ShowEquipmentInfoResponseHaTypeEnum struct {
	ACTIVE  ShowEquipmentInfoResponseHaType
	STANDBY ShowEquipmentInfoResponseHaType
}

func GetShowEquipmentInfoResponseHaTypeEnum() ShowEquipmentInfoResponseHaTypeEnum {
	return ShowEquipmentInfoResponseHaTypeEnum{
		ACTIVE: ShowEquipmentInfoResponseHaType{
			value: "Active",
		},
		STANDBY: ShowEquipmentInfoResponseHaType{
			value: "Standby",
		},
	}
}

func (c ShowEquipmentInfoResponseHaType) Value() string {
	return c.value
}

func (c ShowEquipmentInfoResponseHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEquipmentInfoResponseHaType) UnmarshalJSON(b []byte) error {
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
