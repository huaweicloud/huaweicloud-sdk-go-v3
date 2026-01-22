package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues struct {

	// **参数解释**： Queue名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 对应的Vhost。 **取值范围**： 不涉及。
	Vhost *string `json:"vhost,omitempty"`
}

func (o ShowCeshierarchyRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues", string(data)}, " ")
}
