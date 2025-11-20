package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncNodeRequest Request Object
type SyncNodeRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 节点ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeId string `json:"node_id"`
}

func (o SyncNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncNodeRequest struct{}"
	}

	return strings.Join([]string{"SyncNodeRequest", string(data)}, " ")
}
