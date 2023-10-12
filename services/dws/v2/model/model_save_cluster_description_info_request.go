package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveClusterDescriptionInfoRequest Request Object
type SaveClusterDescriptionInfoRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	Body *ClusterDescriptionInfo `json:"body,omitempty"`
}

func (o SaveClusterDescriptionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveClusterDescriptionInfoRequest struct{}"
	}

	return strings.Join([]string{"SaveClusterDescriptionInfoRequest", string(data)}, " ")
}
