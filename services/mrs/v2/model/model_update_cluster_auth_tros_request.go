package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterAuthTrosRequest Request Object
type UpdateClusterAuthTrosRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *ClusterAuthTros `json:"body,omitempty"`
}

func (o UpdateClusterAuthTrosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterAuthTrosRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterAuthTrosRequest", string(data)}, " ")
}
