package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisasterRecovery struct {

	// **参数解释**： 容灾ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - creating，容灾创建中。 - create_failed，容灾创建失败。 - unstart，容灾未启动。 - starting，容灾启动中。 - start_failed，容灾启动失败。 - running，容灾运行中。 - stopping，容灾停止中。 - stop_failed，容灾停止失败。 - switchovering，灾备切换中。 - abnormal，容灾异常。 - deleting，容灾删除中。 - deleted，容灾已删除。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 容灾类型。 **取值范围**： - az，跨az容灾。 - region，跨region容灾。
	DrType *string `json:"dr_type,omitempty"`

	// **参数解释**： 主集群ID。 **取值范围**： 不涉及。
	PrimaryClusterId *string `json:"primary_cluster_id,omitempty"`

	// **参数解释**： 主集群名称。 **取值范围**： 不涉及。
	PrimaryClusterName *string `json:"primary_cluster_name,omitempty"`

	// **参数解释**： 备集群ID。 **取值范围**： 不涉及。
	StandbyClusterId *string `json:"standby_cluster_id,omitempty"`

	// **参数解释**： 备集群名称。 **取值范围**： 不涉及。
	StandbyClusterName *string `json:"standby_cluster_name,omitempty"`

	// **参数解释**： 主集群角色。 **取值范围**： 不涉及。
	PrimaryClusterRole *string `json:"primary_cluster_role,omitempty"`

	// **参数解释**： 备集群角色。 **取值范围**： 不涉及。
	StandbyClusterRole *string `json:"standby_cluster_role,omitempty"`

	// **参数解释**： 主集群状态。 **取值范围**： - backuping，备份运行中。 - restoring，恢复运行中。 - stopped，集群已停止。 - waiting，容灾运行中。 - abnormal，容灾异常。 - drDeleted，容灾已删除。
	PrimaryClusterStatus *string `json:"primary_cluster_status,omitempty"`

	// **参数解释**： 备集群状态。 **取值范围**： - backuping，备份运行中。 - restoring，恢复运行中。 - stopped，集群已停止。 - waiting，容灾运行中。 - abnormal，容灾异常。 - drDeleted，容灾已删除。
	StandbyClusterStatus *string `json:"standby_cluster_status,omitempty"`

	// **参数解释**： 主集群region。 **取值范围**： 不涉及。
	PrimaryClusterRegion *string `json:"primary_cluster_region,omitempty"`

	// **参数解释**： 备集群region。 **取值范围**： 不涉及。
	StandbyClusterRegion *string `json:"standby_cluster_region,omitempty"`

	// **参数解释**： 主集群项目ID。 **取值范围**： 不涉及。
	PrimaryClusterProjectId *string `json:"primary_cluster_project_id,omitempty"`

	// **参数解释**： 备集群项目ID。 **取值范围**： 不涉及。
	StandbyClusterProjectId *string `json:"standby_cluster_project_id,omitempty"`

	// **参数解释**： 最近同步时间。 **取值范围**： 不涉及。
	LastDisasterTime *string `json:"last_disaster_time,omitempty"`

	// **参数解释**： 启动时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o DisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecovery struct{}"
	}

	return strings.Join([]string{"DisasterRecovery", string(data)}, " ")
}
