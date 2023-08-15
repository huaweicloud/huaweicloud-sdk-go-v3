package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

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
