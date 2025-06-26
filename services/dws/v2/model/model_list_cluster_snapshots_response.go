package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterSnapshotsResponse Response Object
type ListClusterSnapshotsResponse struct {

	// **参数解释**： 快照对象列表。 **取值范围**： 不涉及。
	Snapshots *[]ClusterSnapshots `json:"snapshots,omitempty"`

	// **参数解释**： 项目ID。获取方法请参见[获取项目ID](dws_02_0011.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 快照对象列表总数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterSnapshotsResponse", string(data)}, " ")
}
