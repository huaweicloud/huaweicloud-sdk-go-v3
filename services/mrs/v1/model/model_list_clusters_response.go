package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersResponse Response Object
type ListClustersResponse struct {

	// 集群列表总数。
	ClusterTotal *int32 `json:"clusterTotal,omitempty"`

	// 集群参数。
	Clusters       *[]Cluster `json:"clusters,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersResponse struct{}"
	}

	return strings.Join([]string{"ListClustersResponse", string(data)}, " ")
}
