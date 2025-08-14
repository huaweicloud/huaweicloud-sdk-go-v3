package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentDaemonsetDetailInfoResponse Response Object
type ShowAgentDaemonsetDetailInfoResponse struct {

	// 原始yaml
	YamlContent *string `json:"yaml_content,omitempty"`

	// 节点总数
	NodeNum *int32 `json:"node_num,omitempty"`

	// 容器运行时配置
	RuntimeInfo *[]RuntimeRequestBody `json:"runtime_info,omitempty"`

	// **参数解释**: 集群状态 **约束限制**: 不涉及 **取值范围**: 包含如下： - Available：可用，表示集群处于正常状态。 - Unavailable：不可用，表示集群异常，需手动删除或联系管理员删除。 - ScalingUp：扩容中，表示集群正处于扩容过程中。  - ScalingDown：缩容中，表示集群正处于缩容过程中。 - Creating：创建中，表示集群正处于创建过程中。  - Deleting：删除中，表示集群正处于删除过程中。 - Upgrading：升级中，表示集群正处于升级过程中。  - Resizing：规格变更中，表示集群正处于变更规格中。 - RollingBack：回滚中，表示集群正处于回滚过程中。 - RollbackFailed：回滚异常，表示集群回滚异常，需联系管理员进行回滚重试。  - Empty：集群无任何资源。  **默认取值**: 不涉及
	ClusterStatus *string `json:"cluster_status,omitempty"`

	DsInfo *ClusterInfoResponseDsInfo `json:"ds_info,omitempty"`

	// **参数解释**: 集群ds安装状态 **约束限制**: 不涉及 **取值范围**: 包含如下： - installing：安装中。 - install_success：安装成功。 - install_failed：安装失败。 - partically_success：部分安装成功。 - upgrade_success：升级成功。  - upgrade_failed：升级失败。 - upgrading：升级中。 - none：未安装。  **默认取值**: 不涉及
	InstalledStatus *string `json:"installed_status,omitempty"`

	ScheduleInfo   *CreateDaemonsetRequestBodyScheduleInfo `json:"schedule_info,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ShowAgentDaemonsetDetailInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentDaemonsetDetailInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentDaemonsetDetailInfoResponse", string(data)}, " ")
}
