package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginResponseInfo 插件信息
type PluginResponseInfo struct {

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 私有IP地址
	PrivateIp *string `json:"private_ip,omitempty"`

	// 公有IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 系统类型
	OsType *string `json:"os_type,omitempty"`

	// 插件名称
	PluginName *string `json:"plugin_name,omitempty"`

	// 插件版本
	PluginVersion *string `json:"plugin_version,omitempty"`

	// 插件状态
	PluginStatus *string `json:"plugin_status,omitempty"`

	// 插件升级状态
	UpgradeStatus *string `json:"upgrade_status,omitempty"`
}

func (o PluginResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginResponseInfo struct{}"
	}

	return strings.Join([]string{"PluginResponseInfo", string(data)}, " ")
}
