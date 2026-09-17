package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonMetadata **参数解释**： 插件基本信息，集合类的元素类型，包含一组由不同名称定义的属性。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type AddonMetadata struct {

	// **参数解释**： 插件实例唯一ID标识，创建成功后系统自动生成。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释**： 插件名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件别名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Alias *string `json:"alias,omitempty"`

	// **参数解释**： 插件标签，key/value对格式，接口保留字段，填写不会生效。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Labels map[string]string `json:"labels,omitempty"`

	// **参数解释**： 插件注解，由key/value组成。 **约束限制**： 不涉及 **取值范围**： - 安装时固定值为{\"addon.install/type\":\"install\"} - 升级时固定值为{\"addon.upgrade/type\":\"upgrade\"}  **默认取值**： 不涉及
	Annotations map[string]string `json:"annotations,omitempty"`

	// **参数解释**： 更新时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	// **参数解释**： 创建时间，创建成功后系统自动生成，填写无效。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

func (o AddonMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonMetadata struct{}"
	}

	return strings.Join([]string{"AddonMetadata", string(data)}, " ")
}
