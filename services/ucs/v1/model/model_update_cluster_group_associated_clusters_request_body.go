package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterGroupAssociatedClustersRequestBody struct {

	// 更新容器舰队关联集群信息。
	ClusterIds *[]string `json:"clusterIds,omitempty"`
}

func (o UpdateClusterGroupAssociatedClustersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedClustersRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedClustersRequestBody", string(data)}, " ")
}
