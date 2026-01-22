package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCesHierarchyRespInstanceIds struct {

	// **参数解释**： 实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCesHierarchyRespInstanceIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCesHierarchyRespInstanceIds struct{}"
	}

	return strings.Join([]string{"ShowCesHierarchyRespInstanceIds", string(data)}, " ")
}
