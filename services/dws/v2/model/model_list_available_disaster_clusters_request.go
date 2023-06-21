package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAvailableDisasterClustersRequest struct {

	// 主集群ID
	PrimaryClusterId string `json:"primary_cluster_id"`

	// 主集群规格ID
	PrimarySpecId *string `json:"primary_spec_id,omitempty"`

	// 主集群DN数量
	PrimaryClusterDnNum *string `json:"primary_cluster_dn_num,omitempty"`

	// 备集群所在region
	StandbyRegion *string `json:"standby_region,omitempty"`

	// 备集群项目ID
	StandbyProjectId *string `json:"standby_project_id,omitempty"`

	// 备集群所在AZ
	StandbyAzCode string `json:"standby_az_code"`

	// 容灾类型
	DrType *string `json:"dr_type,omitempty"`

	// 数仓类型
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 数仓版本
	DatastoreVersion *string `json:"datastore_version,omitempty"`
}

func (o ListAvailableDisasterClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableDisasterClustersRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableDisasterClustersRequest", string(data)}, " ")
}
