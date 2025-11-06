package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootClusterRequest Request Object
type RebootClusterRequest struct {

	// 指定待操作的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o RebootClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootClusterRequest struct{}"
	}

	return strings.Join([]string{"RebootClusterRequest", string(data)}, " ")
}
