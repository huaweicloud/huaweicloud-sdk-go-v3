package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperNodeMetadata struct {

	// **参数解释**： 超节点名称 > 命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围1-56位，且不能以中划线(-)结尾。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 超节点ID，资源唯一标识，创建成功后自动生成，填写无效
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 创建时间，创建成功后自动生成，填写无效
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**： 更新时间，创建成功后自动生成，填写无效
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	OwnerReference *HyperNodeMetadataOwnerReference `json:"ownerReference,omitempty"`
}

func (o HyperNodeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperNodeMetadata struct{}"
	}

	return strings.Join([]string{"HyperNodeMetadata", string(data)}, " ")
}
