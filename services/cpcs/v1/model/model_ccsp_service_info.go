package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CcspServiceInfo struct {

	// 当前租户拥有的专属密码服务集群数量
	ClusterNum int32 `json:"cluster_num"`

	// 当前租户拥有的专属密码服务实例数量
	InstanceNum int32 `json:"instance_num"`

	// 当前租户的可创建的专属密码服务实例配额数
	InstanceQuota int32 `json:"instance_quota"`

	InstanceDistribution *InstanceDistribution `json:"instance_distribution"`
}

func (o CcspServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcspServiceInfo struct{}"
	}

	return strings.Join([]string{"CcspServiceInfo", string(data)}, " ")
}
