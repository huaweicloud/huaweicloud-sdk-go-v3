package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputingClusterResponse Response Object
type ListComputingClusterResponse struct {

	// 计算集群列表。
	Clusters *[]ComputingClusterRsp `json:"clusters,omitempty"`

	// 计算集群总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListComputingClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingClusterResponse struct{}"
	}

	return strings.Join([]string{"ListComputingClusterResponse", string(data)}, " ")
}
