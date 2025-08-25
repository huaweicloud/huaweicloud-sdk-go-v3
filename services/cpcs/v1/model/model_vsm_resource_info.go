package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VsmResourceInfo struct {

	// 当前租户拥有的vsm集群数量
	ClusterNum int32 `json:"cluster_num"`

	// 当前租户名下由cpcs代创建和管理的vsm集群数量
	CpcsClusterNum int32 `json:"cpcs_cluster_num"`

	// 当前租户拥有的vsm实例总数
	InstanceNum int32 `json:"instance_num"`

	// 当前租户名下由cpcs代创建和管理的vsm实例总数
	CpcsInstanceNum int32 `json:"cpcs_instance_num"`

	// 当前租户被分配的vsm实例配额数
	InstanceQuota int32 `json:"instance_quota"`
}

func (o VsmResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VsmResourceInfo struct{}"
	}

	return strings.Join([]string{"VsmResourceInfo", string(data)}, " ")
}
