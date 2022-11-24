package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AddSubsetsToGatewayResponseBody struct {

	// 设备ID
	Id *int32 `json:"id,omitempty"`

	// 设备ID（兼容20.0）
	DeviceId *int32 `json:"device_id,omitempty"`

	// 父设备ID
	ParentDeviceId *int32 `json:"parent_device_id,omitempty"`

	Product *ProductReferer `json:"product,omitempty"`

	// 设备名称，支持中文、中文标点符号（）。；，：“”、？《》及英文大小写、数字及英文符号()_,#.?'-@%&!, 长度2-64
	DeviceName *string `json:"device_name,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 设备客户端ID，平台生成的设备唯一标识
	ClientId *string `json:"client_id,omitempty"`

	// 设备物理编号，通常使用MAC或者IMEI号，支持英文大小写，数字，下划线和中划线，长度2-64
	NodeId *string `json:"node_id,omitempty"`

	// 设备状态 0-启用 1-禁用
	Status *AddSubsetsToGatewayResponseBodyStatus `json:"status,omitempty"`

	// 是否在线 0-未连接 1-在线 2-离线
	OnlineStatus *AddSubsetsToGatewayResponseBodyOnlineStatus `json:"online_status,omitempty"`

	// 备注
	Description *string `json:"description,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`
}

func (o AddSubsetsToGatewayResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubsetsToGatewayResponseBody struct{}"
	}

	return strings.Join([]string{"AddSubsetsToGatewayResponseBody", string(data)}, " ")
}

type AddSubsetsToGatewayResponseBodyStatus struct {
	value int32
}

type AddSubsetsToGatewayResponseBodyStatusEnum struct {
	E_0 AddSubsetsToGatewayResponseBodyStatus
	E_1 AddSubsetsToGatewayResponseBodyStatus
}

func GetAddSubsetsToGatewayResponseBodyStatusEnum() AddSubsetsToGatewayResponseBodyStatusEnum {
	return AddSubsetsToGatewayResponseBodyStatusEnum{
		E_0: AddSubsetsToGatewayResponseBodyStatus{
			value: 0,
		}, E_1: AddSubsetsToGatewayResponseBodyStatus{
			value: 1,
		},
	}
}

func (c AddSubsetsToGatewayResponseBodyStatus) Value() int32 {
	return c.value
}

func (c AddSubsetsToGatewayResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddSubsetsToGatewayResponseBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AddSubsetsToGatewayResponseBodyOnlineStatus struct {
	value int32
}

type AddSubsetsToGatewayResponseBodyOnlineStatusEnum struct {
	E_0 AddSubsetsToGatewayResponseBodyOnlineStatus
	E_1 AddSubsetsToGatewayResponseBodyOnlineStatus
	E_2 AddSubsetsToGatewayResponseBodyOnlineStatus
}

func GetAddSubsetsToGatewayResponseBodyOnlineStatusEnum() AddSubsetsToGatewayResponseBodyOnlineStatusEnum {
	return AddSubsetsToGatewayResponseBodyOnlineStatusEnum{
		E_0: AddSubsetsToGatewayResponseBodyOnlineStatus{
			value: 0,
		}, E_1: AddSubsetsToGatewayResponseBodyOnlineStatus{
			value: 1,
		}, E_2: AddSubsetsToGatewayResponseBodyOnlineStatus{
			value: 2,
		},
	}
}

func (c AddSubsetsToGatewayResponseBodyOnlineStatus) Value() int32 {
	return c.value
}

func (c AddSubsetsToGatewayResponseBodyOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddSubsetsToGatewayResponseBodyOnlineStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
