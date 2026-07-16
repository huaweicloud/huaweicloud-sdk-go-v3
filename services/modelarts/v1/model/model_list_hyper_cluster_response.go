package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperClusterResponse Response Object
type ListHyperClusterResponse struct {

	// **参数解释**：Hyper Cluster列表。
	HyperClusters  *[]HyperCluster `json:"hyper_clusters,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListHyperClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperClusterResponse struct{}"
	}

	return strings.Join([]string{"ListHyperClusterResponse", string(data)}, " ")
}
