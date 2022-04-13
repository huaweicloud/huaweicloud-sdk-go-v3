package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartClusterRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`

	Body *CdmStartClusterReq `json:"body,omitempty"`
}

func (o StartClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartClusterRequest struct{}"
	}

	return strings.Join([]string{"StartClusterRequest", string(data)}, " ")
}
