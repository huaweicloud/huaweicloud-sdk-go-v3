package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterPlansRequest Request Object
type ListLogicalClusterPlansRequest struct {

	// 指定集群的ID
	ClusterId string `json:"cluster_id"`
}

func (o ListLogicalClusterPlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterPlansRequest struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterPlansRequest", string(data)}, " ")
}
