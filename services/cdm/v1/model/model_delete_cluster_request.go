package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteClusterRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`

	Body *CdmDeleteClusterReq `json:"body,omitempty"`
}

func (o DeleteClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
