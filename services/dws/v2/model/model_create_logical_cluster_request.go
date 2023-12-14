package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogicalClusterRequest Request Object
type CreateLogicalClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *CreateLogicalClusterRequestBody `json:"body,omitempty"`
}

func (o CreateLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateLogicalClusterRequest", string(data)}, " ")
}
