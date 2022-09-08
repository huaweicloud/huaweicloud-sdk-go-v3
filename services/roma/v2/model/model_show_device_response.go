package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDeviceResponse struct {

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
	Status *ShowDeviceResponseStatus `json:"status,omitempty"`

	// 是否在线 0-未连接 1-在线 2-离线
	OnlineStatus *ShowDeviceResponseOnlineStatus `json:"online_status,omitempty"`

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
	DeviceType *ShowDeviceResponseDeviceType `json:"device_type,omitempty"`

	// 客户端ip
	ClientIp *string `json:"client_ip,omitempty"`

	// 心跳时间
	KeepAlive *string `json:"keep_alive,omitempty"`

	// 最后登录时间
	LastActiveTime *int64 `json:"last_active_time,omitempty"`

	// 设备版本
	Version *string `json:"version,omitempty"`

	// modbus和opcua设备特有,表示设备所属产品的类型 0-普通产品 1-modbus网关产品 2-opcua网关产品
	PluginId *ShowDeviceResponsePluginId `json:"plugin_id,omitempty"`

	// 应用ID
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceResponse", string(data)}, " ")
}

type ShowDeviceResponseStatus struct {
	value int32
}

type ShowDeviceResponseStatusEnum struct {
	E_0 ShowDeviceResponseStatus
	E_1 ShowDeviceResponseStatus
}

func GetShowDeviceResponseStatusEnum() ShowDeviceResponseStatusEnum {
	return ShowDeviceResponseStatusEnum{
		E_0: ShowDeviceResponseStatus{
			value: 0,
		}, E_1: ShowDeviceResponseStatus{
			value: 1,
		},
	}
}

func (c ShowDeviceResponseStatus) Value() int32 {
	return c.value
}

func (c ShowDeviceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeviceResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowDeviceResponseOnlineStatus struct {
	value int32
}

type ShowDeviceResponseOnlineStatusEnum struct {
	E_0 ShowDeviceResponseOnlineStatus
	E_1 ShowDeviceResponseOnlineStatus
	E_2 ShowDeviceResponseOnlineStatus
}

func GetShowDeviceResponseOnlineStatusEnum() ShowDeviceResponseOnlineStatusEnum {
	return ShowDeviceResponseOnlineStatusEnum{
		E_0: ShowDeviceResponseOnlineStatus{
			value: 0,
		}, E_1: ShowDeviceResponseOnlineStatus{
			value: 1,
		}, E_2: ShowDeviceResponseOnlineStatus{
			value: 2,
		},
	}
}

func (c ShowDeviceResponseOnlineStatus) Value() int32 {
	return c.value
}

func (c ShowDeviceResponseOnlineStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeviceResponseOnlineStatus) UnmarshalJSON(b []byte) error {
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

type ShowDeviceResponseDeviceType struct {
	value int32
}

type ShowDeviceResponseDeviceTypeEnum struct {
	E_0 ShowDeviceResponseDeviceType
	E_1 ShowDeviceResponseDeviceType
	E_2 ShowDeviceResponseDeviceType
}

func GetShowDeviceResponseDeviceTypeEnum() ShowDeviceResponseDeviceTypeEnum {
	return ShowDeviceResponseDeviceTypeEnum{
		E_0: ShowDeviceResponseDeviceType{
			value: 0,
		}, E_1: ShowDeviceResponseDeviceType{
			value: 1,
		}, E_2: ShowDeviceResponseDeviceType{
			value: 2,
		},
	}
}

func (c ShowDeviceResponseDeviceType) Value() int32 {
	return c.value
}

func (c ShowDeviceResponseDeviceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeviceResponseDeviceType) UnmarshalJSON(b []byte) error {
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

type ShowDeviceResponsePluginId struct {
	value int32
}

type ShowDeviceResponsePluginIdEnum struct {
	E_0 ShowDeviceResponsePluginId
	E_1 ShowDeviceResponsePluginId
	E_2 ShowDeviceResponsePluginId
}

func GetShowDeviceResponsePluginIdEnum() ShowDeviceResponsePluginIdEnum {
	return ShowDeviceResponsePluginIdEnum{
		E_0: ShowDeviceResponsePluginId{
			value: 0,
		}, E_1: ShowDeviceResponsePluginId{
			value: 1,
		}, E_2: ShowDeviceResponsePluginId{
			value: 2,
		},
	}
}

func (c ShowDeviceResponsePluginId) Value() int32 {
	return c.value
}

func (c ShowDeviceResponsePluginId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeviceResponsePluginId) UnmarshalJSON(b []byte) error {
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
