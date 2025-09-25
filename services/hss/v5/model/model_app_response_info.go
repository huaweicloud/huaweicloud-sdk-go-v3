package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppResponseInfo 软件响应信息
type AppResponseInfo struct {

	// **参数解释**： agent ID **取值范围**： 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 服务器 ID **取值范围**： 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 服务器名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器 IP **取值范围**： 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**： 软件名称 **取值范围**： 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**： 软件版本 **取值范围**： 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 软件安装路径 **取值范围**： 不涉及
	InstallDir *string `json:"install_dir,omitempty"`

	// **参数解释**： 容器 ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释** 最近扫描时间 **取值范围** 不涉及
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`

	// **参数解释** 更新时间 **取值范围** 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AppResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppResponseInfo struct{}"
	}

	return strings.Join([]string{"AppResponseInfo", string(data)}, " ")
}
