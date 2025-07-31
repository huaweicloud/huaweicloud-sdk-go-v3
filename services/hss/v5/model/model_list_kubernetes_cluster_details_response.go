package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKubernetesClusterDetailsResponse Response Object
type ListKubernetesClusterDetailsResponse struct {

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群列表
	ClusterInfoList *[]KubernetesClusterInfo `json:"cluster_info_list,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListKubernetesClusterDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKubernetesClusterDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListKubernetesClusterDetailsResponse", string(data)}, " ")
}
