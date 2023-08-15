package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkClusterRequest Request Object
type ShrinkClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *ClusterShrinkReq `json:"body,omitempty"`
}

func (o ShrinkClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkClusterRequest struct{}"
	}

	return strings.Join([]string{"ShrinkClusterRequest", string(data)}, " ")
}
