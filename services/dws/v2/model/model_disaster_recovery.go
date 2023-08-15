package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisasterRecovery struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 容灾类型
	DrType *string `json:"dr_type,omitempty"`

	// 主集群ID
	PrimaryClusterId *string `json:"primary_cluster_id,omitempty"`

	// 主集群名称
	PrimaryClusterName *string `json:"primary_cluster_name,omitempty"`

	// 备集群ID
	StandbyClusterId *string `json:"standby_cluster_id,omitempty"`

	// 备集群名称
	StandbyClusterName *string `json:"standby_cluster_name,omitempty"`

	// 主集群角色
	PrimaryClusterRole *string `json:"primary_cluster_role,omitempty"`

	// 备集群角色
	StandbyClusterRole *string `json:"standby_cluster_role,omitempty"`

	// 主集群状态
	PrimaryClusterStatus *string `json:"primary_cluster_status,omitempty"`

	// 备集群状态
	StandbyClusterStatus *string `json:"standby_cluster_status,omitempty"`

	// 主集群region
	PrimaryClusterRegion *string `json:"primary_cluster_region,omitempty"`

	// 备集群region
	StandbyClusterRegion *string `json:"standby_cluster_region,omitempty"`

	// 主集群project_id
	PrimaryClusterProjectId *string `json:"primary_cluster_project_id,omitempty"`

	// 备集群project_id
	StandbyClusterProjectId *string `json:"standby_cluster_project_id,omitempty"`

	// 最近同步时间
	LastDisasterTime *string `json:"last_disaster_time,omitempty"`

	// 启动时间
	StartTime *string `json:"start_time,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o DisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecovery struct{}"
	}

	return strings.Join([]string{"DisasterRecovery", string(data)}, " ")
}
