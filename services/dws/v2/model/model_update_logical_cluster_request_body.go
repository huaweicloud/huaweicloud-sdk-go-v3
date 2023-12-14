package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterRequestBody 编辑逻辑集群请求体
type UpdateLogicalClusterRequestBody struct {

	// 逻辑集群编辑环列表信息
	ClusterRings []ClusterRing `json:"cluster_rings"`

	// 模式
	Mode *string `json:"mode,omitempty"`

	// 是否等待销毁
	WaitingForKilling *int32 `json:"waiting_for_killing,omitempty"`
}

func (o UpdateLogicalClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterRequestBody", string(data)}, " ")
}
