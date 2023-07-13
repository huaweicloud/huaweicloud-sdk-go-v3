package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubDevice struct {

	// 设备ID
	Id *int32 `json:"id,omitempty"`

	// 设备ID（兼容20.0）
	DeviceId *int32 `json:"device_id,omitempty"`

	// 父设备ID
	ParentDeviceId *int32 `json:"parent_device_id,omitempty"`

	// 父设备名称
	ParentDeviceName *string `json:"parent_device_name,omitempty"`

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
	Status *SubDeviceStatus `json:"status,omitempty"`

	// 是否在线 0-未连接 1-在线 2-离线
	OnlineStatus *SubDeviceOnlineStatus `json:"online_status,omitempty"`

	// 备注
	Description *string `json:"description,omitempty"`

	Authentication *Authentication `json:"authentication,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 创建时间，timestamp(ms)，使用UTC时区
	CreatedDatetime *int64 `json:"created_datetime,omitempty"`

	// 最后修改时间，timestamp(ms)，使用UTC时区
	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`

	// 设备接入地址
	ConnectAddress *string `json:"connect_address,omitempty"`

	// 设备接入SSL地址
	SslConnectAddress *string `json:"ssl_connect_address,omitempty"`

	// 设备接入IPV6地址
	Ipv6ConnectAddress *string `json:"ipv6_connect_address,omitempty"`

	// 设备接入IPV6 SSL地址
	Ipv6SslConnectAddress *string `json:"ipv6_ssl_connect_address,omitempty"`

	// 应用ID
	AppId *string `json:"app_id,omitempty"`
}

func (o SubDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubDevice struct{}"
	}

	return strings.Join([]string{"SubDevice", string(data)}, " ")
}

type SubDeviceStatus struct {
	value int32
}

type SubDeviceStatusEnum struct {
	E_0 SubDeviceStatus
	E_1 SubDeviceStatus
}

func GetSubDeviceStatusEnum() SubDeviceStatusEnum {
	return SubDeviceStatusEnum{
		E_0: SubDeviceStatus{
			value: 0,
		}, E_1: SubDeviceStatus{
			value: 1,
		},
	}
}

func (c SubDeviceStatus) Value() int32 {
	return c.value
}

func (c SubDeviceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubDeviceStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type SubDeviceOnlineStatus struct {
	value int32
}

type SubDeviceOnlineStatusEnum struct {
	E_0 SubDeviceOnlineStatus
	E_1 SubDeviceOnlineStatus
	E_2 SubDeviceOnlineStatus
}

func GetSubDeviceOnlineStatusEnum() SubDeviceOnlineStatusEnum {
	return SubDeviceOnlineStatusEnum{
		E_0: SubDeviceOnlineStatus{
			value: 0,
		}, E_1: SubDeviceOnlineStatus{
			value: 1,
		}, E_2: SubDeviceOnlineStatus{
			value: 2,
		},
	}
}

func (c SubDeviceOnlineStatus) Value() int32 {
	return c.value
}

func (c SubDeviceOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubDeviceOnlineStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
