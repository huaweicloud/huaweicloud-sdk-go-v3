package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterItemResponseInfo 集群防护策略列表
type ClusterItemResponseInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群命名空间
	ClusterNs *[]string `json:"cluster_ns,omitempty"`

	// 集群标签
	ClusterLabels *[]string `json:"cluster_labels,omitempty"`

	// **参数解释**： 集群防护状态 **取值范围**: - unprotected：未受保护。 - plugin error：插件错误。 - protected with policy：受策略保护。 - deploy policy failed：部署策略失败。 - protected without policy：无策略保护。 - uninstall failed：卸载失败。 - uninstall：卸载。 - installing：正在安装。
	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ClusterItemResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterItemResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterItemResponseInfo", string(data)}, " ")
}
