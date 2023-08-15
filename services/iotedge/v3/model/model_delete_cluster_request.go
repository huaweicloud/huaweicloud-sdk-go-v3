package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterRequest Request Object
type DeleteClusterRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`
}

func (o DeleteClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
