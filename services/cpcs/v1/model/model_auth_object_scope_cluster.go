package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthObjectScopeCluster 集群信息
type AuthObjectScopeCluster struct {

	// 集群id
	Id *string `json:"id,omitempty"`

	// 集群名称
	Name *string `json:"name,omitempty"`
}

func (o AuthObjectScopeCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthObjectScopeCluster struct{}"
	}

	return strings.Join([]string{"AuthObjectScopeCluster", string(data)}, " ")
}
