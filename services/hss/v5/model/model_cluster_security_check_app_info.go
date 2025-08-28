package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterSecurityCheckAppInfo 应用信息
type ClusterSecurityCheckAppInfo struct {

	// **参数解释**： 软件名称 **取值范围**： 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**： 软件版本 **取值范围**： 不涉及
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： 节点名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 私有IP地址 **取值范围**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 公有IP地址 **取值范围**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 更新时间 **取值范围**： 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 最近扫描时间 **取值范围**： 不涉及
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o ClusterSecurityCheckAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckAppInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckAppInfo", string(data)}, " ")
}
