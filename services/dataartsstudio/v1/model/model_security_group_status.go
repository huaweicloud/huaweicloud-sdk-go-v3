package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupStatus 集群认证信息
type SecurityGroupStatus struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 安全组名称
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 风险说明
	GroupDescription *string `json:"group_description,omitempty"`
}

func (o SecurityGroupStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupStatus struct{}"
	}

	return strings.Join([]string{"SecurityGroupStatus", string(data)}, " ")
}
