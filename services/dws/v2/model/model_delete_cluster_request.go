package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteClusterRequest struct {

	// 指定待删除集群的ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *DeleteClusterRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DeleteClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
