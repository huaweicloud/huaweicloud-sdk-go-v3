package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagQuotaRequest Request Object
type ShowTagQuotaRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 是否查询弹性伸缩策略标签
	AutoScalingPolicyTags *bool `json:"auto_scaling_policy_tags,omitempty"`
}

func (o ShowTagQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowTagQuotaRequest", string(data)}, " ")
}
