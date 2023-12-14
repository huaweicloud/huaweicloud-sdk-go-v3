package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConvertToLogicalClusterRequest Request Object
type ConvertToLogicalClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 逻辑集群名称
	Name string `json:"name"`
}

func (o ConvertToLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConvertToLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"ConvertToLogicalClusterRequest", string(data)}, " ")
}
