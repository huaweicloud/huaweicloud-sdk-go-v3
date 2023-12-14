package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogicalClusterRequest Request Object
type EnableLogicalClusterRequest struct {

	// 指定集群的ID
	ClusterId string `json:"cluster_id"`

	Body *EnableLogicalClusterRequestBody `json:"body,omitempty"`
}

func (o EnableLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"EnableLogicalClusterRequest", string(data)}, " ")
}
