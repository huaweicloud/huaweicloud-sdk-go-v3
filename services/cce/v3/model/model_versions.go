package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Versions **参数解释**： 具体插件版本信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type Versions struct {

	// **参数解释**： 插件版本号。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Version string `json:"version"`

	// **参数解释**： 插件安装参数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Input *interface{} `json:"input"`

	// **参数解释**： 是否为稳定版本。 **约束限制**： 不涉及 **取值范围**： - true：稳定版本 - false：非稳定版本  **默认取值**： 不涉及
	Stable bool `json:"stable"`

	// **参数解释**： 供界面使用的翻译信息。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Translate *interface{} `json:"translate"`

	// **参数解释**： 支持集群版本号。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SupportVersions []SupportVersions `json:"supportVersions"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**： 更新时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTimestamp string `json:"updateTimestamp"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}
