package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AliasTarget **参数解释：** 别名记录。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type AliasTarget struct {

	// **参数解释：** 资源服务类型，支持别名记录的服务。 **约束限制：** 不涉及。 **取值范围：** - cloudsite：企业门户 - waf：Web应用防火墙  **默认取值：** 不涉及。
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释：** 对应服务下的域名，由各服务提供。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ResourceDomainName *string `json:"resource_domain_name,omitempty"`
}

func (o AliasTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AliasTarget struct{}"
	}

	return strings.Join([]string{"AliasTarget", string(data)}, " ")
}
