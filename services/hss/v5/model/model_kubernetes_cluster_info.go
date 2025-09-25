package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KubernetesClusterInfo 集群列表对象
type KubernetesClusterInfo struct {

	// id
	Id *string `json:"id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 集群类型 **约束限制**: 不涉及 **取值范围**: - k8s：原生集群 - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群  **默认取值**: 不涉及
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释**: 集群状态(11种) **约束限制**: 不涉及 **取值范围**:   - Available：可用，表示集群处于正常状态。   - Unavailable：不可用，表示集群异常，需手动删除或联系管理员删除。   - ScalingUp：扩容中，表示集群正处于扩容过程中。   - ScalingDown：缩容中，表示集群正处于缩容过程中。   - Creating：创建中，表示集群正处于创建过程中。   - Deleting：删除中，表示集群正处于删除过程中。   - Upgrading：升级中，表示集群正处于升级过程中。   - Resizing：规格变更中，表示集群正处于变更规格中。   - RollingBack：回滚中，表示集群正处于回滚过程中。   - RollbackFailed：回滚异常，表示集群回滚异常，需联系管理员进行回滚重试。   - Empty：集群无任何资源  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// 集群版本
	Version *string `json:"version,omitempty"`

	// 节点总数
	TotalNodesNumber *int32 `json:"total_nodes_number,omitempty"`

	// 正常节点数
	ActiveNodesNumber *int32 `json:"active_nodes_number,omitempty"`

	// 创建时间戳
	CreationTimestamp *int64 `json:"creation_timestamp,omitempty"`

	// 集群下已安装agent节点数
	AgentInstalledNum *int32 `json:"agent_installed_num,omitempty"`

	// 集群下安装失败节点数
	AgentInstallFailedNum *int32 `json:"agent_install_failed_num,omitempty"`

	// 集群下未安装agent节点数
	AgentNotInstallNum *int32 `json:"agent_not_install_num,omitempty"`

	// 集群下agent-ds安装状态，关联agent相关信息时则需同时考虑last_operate_time时间，包含如下2种。   - NotInstall：未状态。   - Installed：已安装。
	AgentDsInstallStatus *string `json:"agent_ds_install_status,omitempty"`

	// 操作失败原因
	AgentDsFailedReason *string `json:"agent_ds_failed_reason,omitempty"`

	// 最近操作时间戳，daemonset脚本安装操作时间，间隔10分钟以内时，agent仍处于安装中
	LastOperateTimestamp *int64 `json:"last_operate_timestamp,omitempty"`

	// 集群最近一次扫描时间戳
	LastScanTime *int64 `json:"last_scan_time,omitempty"`

	// 集群下系统漏洞个数
	SysVulNum *int32 `json:"sys_vul_num,omitempty"`

	// 集群下应用漏洞个数
	AppVulNum *int32 `json:"app_vul_num,omitempty"`

	// 集群下应急漏洞个数
	EmgVulNum *int32 `json:"emg_vul_num,omitempty"`

	// 集群下风险评估问题个数
	RiskAssessNum *int32 `json:"risk_assess_num,omitempty"`

	// 集群下安全合规问题个数
	SecCompNum *int32 `json:"sec_comp_num,omitempty"`
}

func (o KubernetesClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KubernetesClusterInfo struct{}"
	}

	return strings.Join([]string{"KubernetesClusterInfo", string(data)}, " ")
}
