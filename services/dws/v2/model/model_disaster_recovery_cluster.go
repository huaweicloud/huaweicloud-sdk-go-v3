package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisasterRecoveryCluster **参数解释**： 容灾集群信息。 **取值范围**： 不涉及。
type DisasterRecoveryCluster struct {

	// **参数解释**： 容灾集群信息ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 容灾集群名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 容灾集群所在可用区。 **取值范围**： 不涉及。
	ClusterAz *string `json:"cluster_az,omitempty"`

	// **参数解释**： 容灾集群角色。 **取值范围**： 不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释**： 容灾集群所在region。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 容灾集群状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 容灾进度。 **取值范围**： 不涉及。
	Progress *string `json:"progress,omitempty"`

	// **参数解释**： 上一次容灾时间。 **取值范围**： 不涉及。
	LastSuccessTime *string `json:"last_success_time,omitempty"`

	// **参数解释**： OBS桶名称。 **取值范围**： 不涉及。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`
}

func (o DisasterRecoveryCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryCluster struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryCluster", string(data)}, " ")
}
