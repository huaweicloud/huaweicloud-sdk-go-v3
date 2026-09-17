package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeNodePoolRequest Request Object
type UpgradeNodePoolRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 选择需要同步/升级的节点池 **约束限制**： 不涉及 **取值范围**： - 节点池ID：同步指定节点池中的配置，节点池ID获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 - DefaultPool：升级默认节点池的配置  **默认取值**： 不涉及
	NodepoolId string `json:"nodepool_id"`

	Body *UpgradeNodePool `json:"body,omitempty"`
}

func (o UpgradeNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeNodePoolRequest struct{}"
	}

	return strings.Join([]string{"UpgradeNodePoolRequest", string(data)}, " ")
}
