package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeMetadataOwnerReferences **参数解释**： 属主对象。 **约束限制**： - 创建成功后自动生成，填写无效。 - 创建节点接口返回内容中无该参数  **取值范围**： 不涉及 **默认取值**： 不涉及
type NodeMetadataOwnerReferences struct {

	// **参数解释**： 节点池名称 **约束限制**： 创建成功后自动生成，填写无效。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolName *string `json:"nodepoolName,omitempty"`

	// **参数解释**： 节点池ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 创建成功后自动生成，填写无效。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolID *string `json:"nodepoolID,omitempty"`

	// **参数解释**： 超节点名称。如果节点不属于超节点，此字段不展示。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HyperNodeName *string `json:"hyperNodeName,omitempty"`

	// **参数解释**： 超节点ID。如果节点不属于超节点，此字段不展示。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HyperNodeID *string `json:"hyperNodeID,omitempty"`
}

func (o NodeMetadataOwnerReferences) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeMetadataOwnerReferences struct{}"
	}

	return strings.Join([]string{"NodeMetadataOwnerReferences", string(data)}, " ")
}
