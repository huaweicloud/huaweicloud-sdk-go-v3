package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoScalingPolicyRequest Request Object
type ShowAutoScalingPolicyRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`
}

func (o ShowAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoScalingPolicyRequest", string(data)}, " ")
}
