package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableDisasterClustersRequest Request Object
type ListAvailableDisasterClustersRequest struct {

	// **参数解释**： 主集群ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PrimaryClusterId string `json:"primary_cluster_id"`

	// **参数解释**： 备集群所在AZ。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyAzCode string `json:"standby_az_code"`

	// **参数解释**： 主集群规格ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PrimarySpecId *string `json:"primary_spec_id,omitempty"`

	// **参数解释**： 主集群DN数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PrimaryClusterDnNum *string `json:"primary_cluster_dn_num,omitempty"`

	// **参数解释**： 备集群所在局点。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyRegion *string `json:"standby_region,omitempty"`

	// **参数解释**： 备集群项目ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyProjectId *string `json:"standby_project_id,omitempty"`

	// **参数解释**： 容灾类型。 **约束限制**： 不涉及。 **取值范围**： - az，跨az容灾。 - region，跨region容灾。 **默认取值**： 不涉及。
	DrType *string `json:"dr_type,omitempty"`

	// **参数解释**： 数仓类型。 **约束限制**： 不涉及。 **取值范围**： - dws，dws存算一体。 - dws3.0，dws存算分离。 - hybrid，dws实时数仓。 **默认取值**： 不涉及。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// **参数解释**： 数仓版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`
}

func (o ListAvailableDisasterClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableDisasterClustersRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableDisasterClustersRequest", string(data)}, " ")
}
