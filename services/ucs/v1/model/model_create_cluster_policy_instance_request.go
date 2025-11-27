package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterPolicyInstanceRequest Request Object
type CreateClusterPolicyInstanceRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	Body *UcsConstraintRequest `json:"body,omitempty"`
}

func (o CreateClusterPolicyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterPolicyInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterPolicyInstanceRequest", string(data)}, " ")
}
