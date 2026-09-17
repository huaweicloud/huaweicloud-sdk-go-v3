package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterResponse Response Object
type ShowClusterResponse struct {

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群描述
	Description *string `json:"description,omitempty"`

	// 边缘集群版本
	Version *string `json:"version,omitempty"`

	// 边缘集群状态
	State *string `json:"state,omitempty"`

	// 操作系统
	Os *string `json:"os,omitempty"`

	// 集群架构
	Arch *string `json:"arch,omitempty"`

	License *LicenseInfo `json:"license,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty"`

	// kubernetes版本
	KubernetesVersion *string `json:"kubernetes_version,omitempty"`

	// 集群license状态
	LicenseStatus *string `json:"license_status,omitempty"`

	// 集群地址
	ClusterAddr *string `json:"cluster_addr,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterResponse", string(data)}, " ")
}
