package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateDeviceResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty"`

	// 设备ID
	Id *int32 `json:"id,omitempty"`

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

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 设备状态 0-启用 1-禁用
	Status *UpdateDeviceResponseStatus `json:"status,omitempty"`

	// 是否在线 0-未连接 1-在线 2-离线
	OnlineStatus *UpdateDeviceResponseOnlineStatus `json:"online_status,omitempty"`

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

	// 最后登录时间
	LastLoginDatetime *int64 `json:"last_login_datetime,omitempty"`

	// 节点类型 0-直连 1-网关 2-子设备
	NodeType *int32 `json:"node_type,omitempty"`

	// 设备类型<br>0-普通设备（无子设备也无父设备）<br>1-网关设备(可挂载子设备)<br>2-子设备(归属于某个网关设备)
	DeviceType *UpdateDeviceResponseDeviceType `json:"device_type,omitempty"`

	// 客户端ip
	ClientIp *string `json:"client_ip,omitempty"`

	// 心跳时间
	KeepAlive *string `json:"keep_alive,omitempty"`

	// 最后登录时间
	LastActiveTime *int64 `json:"last_active_time,omitempty"`

	// 设备版本
	Version *string `json:"version,omitempty"`

	// modbus和opcua设备特有,表示设备所属产品的类型 0-普通产品 1-modbus网关产品 2-opcua网关产品
	PluginId *UpdateDeviceResponsePluginId `json:"plugin_id,omitempty"`

	// 应用ID
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceResponse", string(data)}, " ")
}

type UpdateDeviceResponseStatus struct {
	value int32
}

type UpdateDeviceResponseStatusEnum struct {
	E_0 UpdateDeviceResponseStatus
	E_1 UpdateDeviceResponseStatus
}

func GetUpdateDeviceResponseStatusEnum() UpdateDeviceResponseStatusEnum {
	return UpdateDeviceResponseStatusEnum{
		E_0: UpdateDeviceResponseStatus{
			value: 0,
		}, E_1: UpdateDeviceResponseStatus{
			value: 1,
		},
	}
}

func (c UpdateDeviceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDeviceResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateDeviceResponseOnlineStatus struct {
	value int32
}

type UpdateDeviceResponseOnlineStatusEnum struct {
	E_0 UpdateDeviceResponseOnlineStatus
	E_1 UpdateDeviceResponseOnlineStatus
	E_2 UpdateDeviceResponseOnlineStatus
}

func GetUpdateDeviceResponseOnlineStatusEnum() UpdateDeviceResponseOnlineStatusEnum {
	return UpdateDeviceResponseOnlineStatusEnum{
		E_0: UpdateDeviceResponseOnlineStatus{
			value: 0,
		}, E_1: UpdateDeviceResponseOnlineStatus{
			value: 1,
		}, E_2: UpdateDeviceResponseOnlineStatus{
			value: 2,
		},
	}
}

func (c UpdateDeviceResponseOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDeviceResponseOnlineStatus) UnmarshalJSON(b []byte) error {
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

type UpdateDeviceResponseDeviceType struct {
	value int32
}

type UpdateDeviceResponseDeviceTypeEnum struct {
	E_0 UpdateDeviceResponseDeviceType
	E_1 UpdateDeviceResponseDeviceType
	E_2 UpdateDeviceResponseDeviceType
}

func GetUpdateDeviceResponseDeviceTypeEnum() UpdateDeviceResponseDeviceTypeEnum {
	return UpdateDeviceResponseDeviceTypeEnum{
		E_0: UpdateDeviceResponseDeviceType{
			value: 0,
		}, E_1: UpdateDeviceResponseDeviceType{
			value: 1,
		}, E_2: UpdateDeviceResponseDeviceType{
			value: 2,
		},
	}
}

func (c UpdateDeviceResponseDeviceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDeviceResponseDeviceType) UnmarshalJSON(b []byte) error {
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

type UpdateDeviceResponsePluginId struct {
	value int32
}

type UpdateDeviceResponsePluginIdEnum struct {
	E_0 UpdateDeviceResponsePluginId
	E_1 UpdateDeviceResponsePluginId
	E_2 UpdateDeviceResponsePluginId
}

func GetUpdateDeviceResponsePluginIdEnum() UpdateDeviceResponsePluginIdEnum {
	return UpdateDeviceResponsePluginIdEnum{
		E_0: UpdateDeviceResponsePluginId{
			value: 0,
		}, E_1: UpdateDeviceResponsePluginId{
			value: 1,
		}, E_2: UpdateDeviceResponsePluginId{
			value: 2,
		},
	}
}

func (c UpdateDeviceResponsePluginId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDeviceResponsePluginId) UnmarshalJSON(b []byte) error {
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
