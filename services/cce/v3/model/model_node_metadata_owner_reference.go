package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeMetadataOwnerReference **参数解释**： 属主对象。 **约束限制**： - 创建成功后自动生成，填写无效。 - 创建节点接口返回内容中无该参数  **取值范围**： 不涉及 **默认取值**： 不涉及
type NodeMetadataOwnerReference struct {

	// **参数解释**： 节点池名称 **约束限制**： 创建成功后自动生成，填写无效。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolName *string `json:"nodepoolName,omitempty"`

	// **参数解释**： 节点池UID **约束限制**： 创建成功后自动生成，填写无效。 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodepoolID *string `json:"nodepoolID,omitempty"`
}

func (o NodeMetadataOwnerReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeMetadataOwnerReference struct{}"
	}

	return strings.Join([]string{"NodeMetadataOwnerReference", string(data)}, " ")
}
