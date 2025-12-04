package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClusterGroupPolicyRequest Request Object
type DisableClusterGroupPolicyRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`
}

func (o DisableClusterGroupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClusterGroupPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisableClusterGroupPolicyRequest", string(data)}, " ")
}
