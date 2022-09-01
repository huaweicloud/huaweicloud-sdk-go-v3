package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClusterDetailRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`
}

func (o ShowClusterDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailRequest", string(data)}, " ")
}
