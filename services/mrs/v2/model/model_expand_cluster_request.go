package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandClusterRequest Request Object
type ExpandClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *ExpandParam `json:"body,omitempty"`
}

func (o ExpandClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandClusterRequest struct{}"
	}

	return strings.Join([]string{"ExpandClusterRequest", string(data)}, " ")
}
