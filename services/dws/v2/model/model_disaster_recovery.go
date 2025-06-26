package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisasterRecovery struct {

	// **参数解释**： 容灾ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 容灾类型。 **取值范围**： 不涉及。
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

	// **参数解释**： 主集群状态。 **取值范围**： 不涉及。
	PrimaryClusterStatus *string `json:"primary_cluster_status,omitempty"`

	// **参数解释**： 备集群状态。 **取值范围**： 不涉及。
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
