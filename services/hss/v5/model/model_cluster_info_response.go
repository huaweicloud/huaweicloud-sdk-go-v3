package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterInfoResponse struct {

	// **参数解释** 是否最新版本 **取值范围** - true：是。 - false：否。
	LatestVersion *bool `json:"latest_version,omitempty"`

	// **参数解释** 集群agent版本 **取值范围** 字符长度0-32位
	AgentVersion *string `json:"agent_version,omitempty"`

	// **参数解释** 集群名称 **取值范围** 字符长度0-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释** 集群id **取值范围** 字符长度1-256位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释** 命名空间 **取值范围** 字符长度1-32位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释** 集群类型 **取值范围** 字符长度1-32位
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释** 节点总数 **取值范围** 取值0-65535
	NodeNum *int32 `json:"node_num,omitempty"`

	DsInfo *ClusterInfoResponseDsInfo `json:"ds_info,omitempty"`

	// **参数解释**: 集群状态 **约束限制**: 不涉及 **取值范围**: 包含如下： - Available：可用，表示集群处于正常状态。 - Unavailable：不可用，表示集群异常，需手动删除或联系管理员删除。 - ScalingUp：扩容中，表示集群正处于扩容过程中。 - ScalingDown：缩容中，表示集群正处于缩容过程中。 - Creating：创建中，表示集群正处于创建过程中。 - Deleting：删除中，表示集群正处于删除过程中。 - Upgrading：升级中，表示集群正处于升级过程中。 - Resizing：规格变更中，表示集群正处于变更规格中。 - RollingBack：回滚中，表示集群正处于回滚过程中。 - RollbackFailed：回滚异常，表示集群回滚异常，需联系管理员进行回滚重试。 - Empty：集群无任何资源。  **默认取值**: 不涉及
	ClusterStatus *string `json:"cluster_status,omitempty"`

	// **参数解释**: 集群ds安装状态 **约束限制**: 不涉及 **取值范围**: 包含如下： - installing：安装中。 - install_success：安装成功。 - install_failed：安装失败。 - partially_success：部分安装成功。 - upgrade_success：升级成功。 - upgrade_failed：升级失败。 - upgrading：升级中。 - none：未安装。  **默认取值**: 不涉及
	InstalledStatus *string `json:"installed_status,omitempty"`

	// **参数解释**： 集群anp接入状态 **约束限制**： 不涉及 **取值范围**： - not_connect：未连接。 - connect_success：连接成功。 - connect_fail：连接失败。 - connect_disruption：连接中断。  **默认取值**： 不涉及
	AccessStatus *string `json:"access_status,omitempty"`

	// **参数解释**： 集群anp与ds的组合状态 **约束限制**： 不涉及 **取值范围**： - accessing：接入中。 - access_error：接入异常。 - running：运行中。 - run_error：运行异常。 - upgrading：升级中。 - upgrade_error：升级失败。  **默认取值**： 不涉及
	CombinedStatus *string `json:"combined_status,omitempty"`

	// **参数解释** 集群插件接入失败的原因 **取值范围** 字符长度1-256位
	FailedMessage *string `json:"failed_message,omitempty"`

	// **参数解释**： 集群日志的接入状态 **约束限制**： 不涉及 **取值范围**： - success：接入成功。 - partial_success：部分接入。  **默认取值**： 不涉及
	ClusterLogStatus *string `json:"cluster_log_status,omitempty"`

	// **参数解释**： 调用服务，标识cce免费体检报告 **约束限制**： 不涉及 **取值范围**： - hss：hss服务。 - cce：cce服务。  **默认取值**： 不涉及
	InvokedService *string `json:"invoked_service,omitempty"`

	RegistryInfo *ClusterInfoResponseRegistryInfo `json:"registry_info,omitempty"`
}

func (o ClusterInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfoResponse struct{}"
	}

	return strings.Join([]string{"ClusterInfoResponse", string(data)}, " ")
}
