package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelReadonlyClusterRequest Request Object
type CancelReadonlyClusterRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o CancelReadonlyClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelReadonlyClusterRequest struct{}"
	}

	return strings.Join([]string{"CancelReadonlyClusterRequest", string(data)}, " ")
}
