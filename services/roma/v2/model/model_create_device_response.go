package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CreateDeviceResponse struct {
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

	Status *CreateDeviceResponseStatus `json:"status,omitempty"`
	// 是否在线 0-未连接 1-在线 2-离线

	OnlineStatus *CreateDeviceResponseOnlineStatus `json:"online_status,omitempty"`
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

	DeviceType *CreateDeviceResponseDeviceType `json:"device_type,omitempty"`
	// 客户端ip

	ClientIp *string `json:"client_ip,omitempty"`
	// 心跳时间

	KeepAlive *string `json:"keep_alive,omitempty"`
	// 最后登录时间

	LastActiveTime *int64 `json:"last_active_time,omitempty"`
	// 设备版本

	Version *string `json:"version,omitempty"`
	// modbus和opcua设备特有,表示设备所属产品的类型 0-普通产品 1-modbus网关产品 2-opcua网关产品

	PluginId *CreateDeviceResponsePluginId `json:"plugin_id,omitempty"`
	// 应用ID

	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeviceResponse struct{}"
	}

	return strings.Join([]string{"CreateDeviceResponse", string(data)}, " ")
}

type CreateDeviceResponseStatus struct {
	value int32
}

type CreateDeviceResponseStatusEnum struct {
	E_0 CreateDeviceResponseStatus
	E_1 CreateDeviceResponseStatus
}

func GetCreateDeviceResponseStatusEnum() CreateDeviceResponseStatusEnum {
	return CreateDeviceResponseStatusEnum{
		E_0: CreateDeviceResponseStatus{
			value: 0,
		}, E_1: CreateDeviceResponseStatus{
			value: 1,
		},
	}
}

func (c CreateDeviceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeviceResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateDeviceResponseOnlineStatus struct {
	value int32
}

type CreateDeviceResponseOnlineStatusEnum struct {
	E_0 CreateDeviceResponseOnlineStatus
	E_1 CreateDeviceResponseOnlineStatus
	E_2 CreateDeviceResponseOnlineStatus
}

func GetCreateDeviceResponseOnlineStatusEnum() CreateDeviceResponseOnlineStatusEnum {
	return CreateDeviceResponseOnlineStatusEnum{
		E_0: CreateDeviceResponseOnlineStatus{
			value: 0,
		}, E_1: CreateDeviceResponseOnlineStatus{
			value: 1,
		}, E_2: CreateDeviceResponseOnlineStatus{
			value: 2,
		},
	}
}

func (c CreateDeviceResponseOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeviceResponseOnlineStatus) UnmarshalJSON(b []byte) error {
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

type CreateDeviceResponseDeviceType struct {
	value int32
}

type CreateDeviceResponseDeviceTypeEnum struct {
	E_0 CreateDeviceResponseDeviceType
	E_1 CreateDeviceResponseDeviceType
	E_2 CreateDeviceResponseDeviceType
}

func GetCreateDeviceResponseDeviceTypeEnum() CreateDeviceResponseDeviceTypeEnum {
	return CreateDeviceResponseDeviceTypeEnum{
		E_0: CreateDeviceResponseDeviceType{
			value: 0,
		}, E_1: CreateDeviceResponseDeviceType{
			value: 1,
		}, E_2: CreateDeviceResponseDeviceType{
			value: 2,
		},
	}
}

func (c CreateDeviceResponseDeviceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeviceResponseDeviceType) UnmarshalJSON(b []byte) error {
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

type CreateDeviceResponsePluginId struct {
	value int32
}

type CreateDeviceResponsePluginIdEnum struct {
	E_0 CreateDeviceResponsePluginId
	E_1 CreateDeviceResponsePluginId
	E_2 CreateDeviceResponsePluginId
}

func GetCreateDeviceResponsePluginIdEnum() CreateDeviceResponsePluginIdEnum {
	return CreateDeviceResponsePluginIdEnum{
		E_0: CreateDeviceResponsePluginId{
			value: 0,
		}, E_1: CreateDeviceResponsePluginId{
			value: 1,
		}, E_2: CreateDeviceResponsePluginId{
			value: 2,
		},
	}
}

func (c CreateDeviceResponsePluginId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDeviceResponsePluginId) UnmarshalJSON(b []byte) error {
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
