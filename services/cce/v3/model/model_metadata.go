package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metadata **参数解释**： 资源描述基本信息，集合类的元素类型，包含一组由不同名称定义的属性。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type Metadata struct {

	// **参数解释**： 唯一id标识 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 资源名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源标签，key/value对格式，接口保留字段，填写不会生效 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Labels map[string]string `json:"labels,omitempty"`

	// **参数解释**： 资源注解，由key/value组成 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Annotations map[string]string `json:"annotations,omitempty"`

	// **参数解释**： 更新时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	// **参数解释**： 创建时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
