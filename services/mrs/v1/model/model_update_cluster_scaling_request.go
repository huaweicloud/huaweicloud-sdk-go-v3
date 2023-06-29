package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterScalingRequest Request Object
type UpdateClusterScalingRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`

	Body *ClusterScalingReq `json:"body,omitempty"`
}

func (o UpdateClusterScalingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterScalingRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterScalingRequest", string(data)}, " ")
}
