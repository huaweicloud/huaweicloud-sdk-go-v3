package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScalingPolicyRequest Request Object
type CreateScalingPolicyRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`

	Body *AutoScalingPolicyReqV11 `json:"body,omitempty"`
}

func (o CreateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyRequest", string(data)}, " ")
}
