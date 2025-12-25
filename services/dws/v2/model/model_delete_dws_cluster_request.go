package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDwsClusterRequest Request Object
type DeleteDwsClusterRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 集群需要保留的快照数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	KeepLastManualBackup *string `json:"keep_last_manual_backup,omitempty"`

	// **参数解释**： 集群是否释放弹性公网IP，默认是NO_RELEASE，不释放绑定的弹性公网IP。 **约束限制**： 不涉及。 **取值范围**： - NO_RELEASE：不释放绑定的弹性公网IP； - RELEASE_BINDING：释放绑定的弹性公网IP；  **默认取值**： NO_RELEASE
	ReleaseEipType *string `json:"release_eip_type,omitempty"`
}

func (o DeleteDwsClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDwsClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteDwsClusterRequest", string(data)}, " ")
}
