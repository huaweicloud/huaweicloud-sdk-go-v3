package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterVolumesResponse Response Object
type ListLogicalClusterVolumesResponse struct {

	// **参数解释**： 逻辑集群磁盘信息列表。 **取值范围**： 不涉及。
	Volumes *[]LogicalClusterVolume `json:"volumes,omitempty"`

	// **参数解释**： 逻辑集群磁盘总数。 **取值范围**： 不涉及。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLogicalClusterVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterVolumesResponse", string(data)}, " ")
}
