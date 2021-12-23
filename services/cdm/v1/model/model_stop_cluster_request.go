package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopClusterRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`

	Body *CdmStopClusterReq `json:"body,omitempty"`
}

func (o StopClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopClusterRequest struct{}"
	}

	return strings.Join([]string{"StopClusterRequest", string(data)}, " ")
}
