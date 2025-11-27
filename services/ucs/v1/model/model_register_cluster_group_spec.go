package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterClusterGroupSpec struct {

	// 关联的集群id
	ClusterIds *[]string `json:"clusterIds,omitempty"`

	// 容器舰队描述信息
	Description *string `json:"description,omitempty"`
}

func (o RegisterClusterGroupSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterGroupSpec struct{}"
	}

	return strings.Join([]string{"RegisterClusterGroupSpec", string(data)}, " ")
}
