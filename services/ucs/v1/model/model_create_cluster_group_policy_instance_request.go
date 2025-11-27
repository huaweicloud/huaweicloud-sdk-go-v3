package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterGroupPolicyInstanceRequest Request Object
type CreateClusterGroupPolicyInstanceRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *UcsConstraintRequest `json:"body,omitempty"`
}

func (o CreateClusterGroupPolicyInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterGroupPolicyInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterGroupPolicyInstanceRequest", string(data)}, " ")
}
