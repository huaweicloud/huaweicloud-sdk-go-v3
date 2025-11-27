package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableClusterPolicyRequest Request Object
type EnableClusterPolicyRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	// 重试启动策略中心
	Retry *string `json:"retry,omitempty"`
}

func (o EnableClusterPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableClusterPolicyRequest struct{}"
	}

	return strings.Join([]string{"EnableClusterPolicyRequest", string(data)}, " ")
}
