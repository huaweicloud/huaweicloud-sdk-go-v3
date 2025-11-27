package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyUnitedWorkloadRequest Request Object
type ShowProxyUnitedWorkloadRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	// 工作负载的类型
	Kind string `json:"kind"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 工作负载的名称
	Name *string `json:"name,omitempty"`
}

func (o ShowProxyUnitedWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyUnitedWorkloadRequest struct{}"
	}

	return strings.Join([]string{"ShowProxyUnitedWorkloadRequest", string(data)}, " ")
}
