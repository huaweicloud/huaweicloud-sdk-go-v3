package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveNodeRequest Request Object
type RemoveNodeRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 移除节点时是否解绑节点默认安全组。 **约束限制**： 不涉及 **取值范围**： - false：移除节点时保留节点默认安全组 - true：移除节点时解绑节点默认安全组  **默认取值**： false
	RemoveNodeSystemSecurityGroup *bool `json:"removeNodeSystemSecurityGroup,omitempty"`

	Body *RemoveNodesTask `json:"body,omitempty"`
}

func (o RemoveNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodeRequest struct{}"
	}

	return strings.Join([]string{"RemoveNodeRequest", string(data)}, " ")
}
