package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDisasterRecoverRequest Request Object
type ListDisasterRecoverRequest struct {

	// **参数解释**： 主集群ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PrimaryClusterId *string `json:"primary_cluster_id,omitempty"`

	// **参数解释**： 备集群ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StandbyClusterId *string `json:"standby_cluster_id,omitempty"`

	// **参数解释**： 容灾ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o ListDisasterRecoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoverRequest struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoverRequest", string(data)}, " ")
}
