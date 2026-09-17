package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HyperNodeSpec **参数解释**： 超节点的配置详情，超节点只包含基本的资源规格属性，其他配置从所属节点池继承。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type HyperNodeSpec struct {

	// **参数解释**： 超节点规格 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**： 所属节点池ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolID *string `json:"nodepoolID,omitempty"`

	NodeTemplate *NodeTemplateInHyperNode `json:"nodeTemplate,omitempty"`

	// **参数解释**： 付费方式 **约束限制**： 不涉及 **取值范围**： - prepaid：预付费，即包年包月； - postpaid：后付费，即按需付费；  **默认取值**： 不涉及
	ChargeMode *string `json:"chargeMode,omitempty"`
}

func (o HyperNodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeSpec struct{}"
	}

	return strings.Join([]string{"HyperNodeSpec", string(data)}, " ")
}
