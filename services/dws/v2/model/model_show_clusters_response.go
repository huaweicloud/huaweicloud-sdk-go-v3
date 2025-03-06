package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClustersResponse Response Object
type ShowClustersResponse struct {

	// 集群对象列表。
	Clusters *[]ClusterInfo `json:"clusters,omitempty"`

	// 集群对象列表总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClustersResponse struct{}"
	}

	return strings.Join([]string{"ShowClustersResponse", string(data)}, " ")
}
