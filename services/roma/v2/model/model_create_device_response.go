package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeviceResponse Response Object
type CreateDeviceResponse struct {

	// 权限
	Permissions *[]string `json:"permissions,omitempty"`

	// 设备ID
	Id *int32 `json:"id,omitempty"`

	// 设备ID（兼容20.0）
	DeviceId *int32 `json:"device_id,omitempty"`

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
	Status *int32 `json:"status,omitempty"`

	// 是否在线 0-未连接 1-在线 2-离线
	OnlineStatus *int32 `json:"online_status,omitempty"`

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

	// 节点类型 0-直连 1-网关 2-子设备
	NodeType *int32 `json:"node_type,omitempty"`

	// 设备类型<br>0-普通设备（无子设备也无父设备）<br>1-网关设备(可挂载子设备)<br>2-子设备(归属于某个网关设备)
	DeviceType *int32 `json:"device_type,omitempty"`

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
