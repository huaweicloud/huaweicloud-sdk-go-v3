package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClusterSnapshotsResponse struct {

	// 快照对象列表。
	Snapshots *[]ClusterSnapshots `json:"snapshots,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 快照对象列表总数
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
