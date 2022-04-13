package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNodePoolRequest struct {
	// 集群 ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。

	ClusterId string `json:"cluster_id"`
	// 节点池ID

	NodepoolId string `json:"nodepool_id"`
	// 集群状态兼容Error参数，用于API平滑切换。 兼容场景下，errorStatus为空则屏蔽Error状态为Deleting状态。

	ErrorStatus *string `json:"errorStatus,omitempty"`
}

func (o ShowNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolRequest struct{}"
	}

	return strings.Join([]string{"ShowNodePoolRequest", string(data)}, " ")
}
