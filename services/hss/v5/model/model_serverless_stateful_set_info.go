package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerlessStatefulSetInfo 有状态工作负载基本信息
type ServerlessStatefulSetInfo struct {

	// 有状态工作负载名称
	Name *string `json:"name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 所属集群
	ClusterName *string `json:"cluster_name,omitempty"`

	// 状态，包含以下几种 -Running：正常运行 -Failed：存在异常
	Status *string `json:"status,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 实例个数
	PodsNum *int32 `json:"pods_num,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 标签
	MatchLabels *[]LabelInfo `json:"match_labels,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 工作负载下已安装agent实例数
	AgentInstalledNum *int32 `json:"agent_installed_num,omitempty"`

	// 工作负载下安装失败实例数
	AgentInstallFailedNum *int32 `json:"agent_install_failed_num,omitempty"`

	// 工作负载下未安装agent实例数
	AgentNotInstallNum *int32 `json:"agent_not_install_num,omitempty"`
}

func (o ServerlessStatefulSetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerlessStatefulSetInfo struct{}"
	}

	return strings.Join([]string{"ServerlessStatefulSetInfo", string(data)}, " ")
}
