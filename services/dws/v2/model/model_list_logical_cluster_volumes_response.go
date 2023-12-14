package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterVolumesResponse Response Object
type ListLogicalClusterVolumesResponse struct {

	// 逻辑集群磁盘信息列表
	Volumes *[]LogicalClusterVolume `json:"volumes,omitempty"`

	// 逻辑集群磁盘总数
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
