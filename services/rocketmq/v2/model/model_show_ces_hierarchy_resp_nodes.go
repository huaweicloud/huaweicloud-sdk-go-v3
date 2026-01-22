package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCesHierarchyRespNodes struct {

	// **参数解释**： 节点名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCesHierarchyRespNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyRespNodes struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyRespNodes", string(data)}, " ")
}
