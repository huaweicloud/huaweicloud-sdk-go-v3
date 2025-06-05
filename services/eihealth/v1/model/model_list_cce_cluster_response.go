package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterResponse Response Object
type ListCceClusterResponse struct {

	// CCE集群列表。
	Clusters *[]CceClusterRsp `json:"clusters,omitempty"`

	// CCE集群集群总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCceClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterResponse struct{}"
	}

	return strings.Join([]string{"ListCceClusterResponse", string(data)}, " ")
}
