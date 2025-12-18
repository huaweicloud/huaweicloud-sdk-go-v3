package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateToNodePoolRequest Request Object
type MigrateToNodePoolRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// 节点池ID
	NodepoolId string `json:"nodepool_id"`

	Body *MigrateNodesToNodePoolList `json:"body,omitempty"`
}

func (o MigrateToNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateToNodePoolRequest struct{}"
	}

	return strings.Join([]string{"MigrateToNodePoolRequest", string(data)}, " ")
}
