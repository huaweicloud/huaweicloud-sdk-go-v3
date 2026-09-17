package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HyperNodeMetadataOwnerReference **参数解释**： 属主对象 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type HyperNodeMetadataOwnerReference struct {

	// **参数解释**： 节点池名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolName *string `json:"nodepoolName,omitempty"`

	// **参数解释**： 节点池ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolID *string `json:"nodepoolID,omitempty"`
}

func (o HyperNodeMetadataOwnerReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeMetadataOwnerReference struct{}"
	}

	return strings.Join([]string{"HyperNodeMetadataOwnerReference", string(data)}, " ")
}
