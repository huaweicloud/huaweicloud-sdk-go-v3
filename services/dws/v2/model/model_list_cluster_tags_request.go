package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterTagsRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterTagsRequest", string(data)}, " ")
}
