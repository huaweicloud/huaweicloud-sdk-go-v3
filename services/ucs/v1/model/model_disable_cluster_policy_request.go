package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClusterPolicyRequest Request Object
type DisableClusterPolicyRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`
}

func (o DisableClusterPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClusterPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisableClusterPolicyRequest", string(data)}, " ")
}
