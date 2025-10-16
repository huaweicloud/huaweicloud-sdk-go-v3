package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterPortRequest Request Object
type ListClusterPortRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterPortRequest struct{}"
	}

	return strings.Join([]string{"ListClusterPortRequest", string(data)}, " ")
}
