package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterRsp 最终租户计算集群详情。
type CceClusterRsp struct {

	// 计算集群ID。
	Id *string `json:"id,omitempty"`

	// 计算集群名称。
	Name *string `json:"name,omitempty"`

	// 计算集群状态。
	Status *string `json:"status,omitempty"`

	// 计算集群版本。
	Version *string `json:"version,omitempty"`

	// 计算集群规格。
	Flavor *string `json:"flavor,omitempty"`

	// 计算集群类别。CCE代表CCE集群，Turbo代表CCE Turbo集群。
	Category *string `json:"category,omitempty"`
}

func (o CceClusterRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterRsp struct{}"
	}

	return strings.Join([]string{"CceClusterRsp", string(data)}, " ")
}
