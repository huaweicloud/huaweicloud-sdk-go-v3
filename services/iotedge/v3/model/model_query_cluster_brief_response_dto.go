package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryClusterBriefResponseDto struct {

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

	// 是否可升级
	IsUpgradeable *bool `json:"is_upgradeable,omitempty"`

	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty"`

	// 集群地址
	ClusterAddr *string `json:"cluster_addr,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryClusterBriefResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryClusterBriefResponseDto struct{}"
	}

	return strings.Join([]string{"QueryClusterBriefResponseDto", string(data)}, " ")
}
