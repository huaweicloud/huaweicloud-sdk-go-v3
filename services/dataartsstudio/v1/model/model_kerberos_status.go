package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KerberosStatus 集群认证信息
type KerberosStatus struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 是否开启了kerberos认证,true:开启，false:未开启
	Status *bool `json:"status,omitempty"`
}

func (o KerberosStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KerberosStatus struct{}"
	}

	return strings.Join([]string{"KerberosStatus", string(data)}, " ")
}
