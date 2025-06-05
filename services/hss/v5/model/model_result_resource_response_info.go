package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultResourceResponseInfo 资源信息
type ResultResourceResponseInfo struct {

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// Agent ID
	AgentId *string `json:"agent_id,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 服务器状态，包含如下4种。   - ACTIVE ：运行中。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// Agent状态，包含如下5种。   - installed ：已安装。   - not_installed ：未安装。   - online ：在线。   - offline ：离线。   - install_failed ：安装失败。   - installing ：安装中。
	AgentStatus *string `json:"agent_status,omitempty"`

	// 防护状态，包含如下2种。 - closed ：未防护。 - opened ：防护中。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// 操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// 操作系统版本
	OsVersion *string `json:"os_version,omitempty"`
}

func (o ResultResourceResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultResourceResponseInfo struct{}"
	}

	return strings.Join([]string{"ResultResourceResponseInfo", string(data)}, " ")
}
