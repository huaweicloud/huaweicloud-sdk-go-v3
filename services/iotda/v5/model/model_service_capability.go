package model

import (
	"encoding/json"

	"strings"
)

// ServiceCapability结构体。
type ServiceCapability struct {
	// 设备的服务ID。

	ServiceId string `json:"service_id"`
	// 设备的服务类型。

	ServiceType string `json:"service_type"`
	// 设备服务支持的属性列表。

	Properties *[]ServiceProperty `json:"properties,omitempty"`
	// 设备服务支持的命令列表。

	Commands *[]ServiceCommand `json:"commands,omitempty"`
	// 设备服务支持的事件列表。

	Events *[]ServiceEvent `json:"events,omitempty"`
	// 设备服务的描述信息。

	Description *string `json:"description,omitempty"`
	// 指定设备服务是否必选。Master（主服务）, Mandatory（必选服务）, Optional（可选服务），目前本字段为非功能性字段，仅起到标识作用。默认为Optional（可选服务）。

	Option *string `json:"option,omitempty"`
}

func (o ServiceCapability) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceCapability struct{}"
	}

	return strings.Join([]string{"ServiceCapability", string(data)}, " ")
}
