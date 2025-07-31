package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyCicdConfigurationRequestBody 修改CI/CD配置
type ModifyCicdConfigurationRequestBody struct {

	// **参数解释** 漏洞白名单 **约束限制** 不涉及 **取值范围** 白名单数量范围0-1000  **默认取值** 不涉及
	VulnerabilityWhitelist *[]string `json:"vulnerability_whitelist,omitempty"`

	// **参数解释** 漏洞黑名单 **约束限制** 不涉及 **取值范围** 黑名单数量范围0-1000  **默认取值** 不涉及
	VulnerabilityBlocklist *[]string `json:"vulnerability_blocklist,omitempty"`

	// **参数解释** 镜像白名单 **约束限制** 不涉及 **取值范围** 白名单数量范围0-1000  **默认取值** 不涉及
	ImageWhitelist *[]string `json:"image_whitelist,omitempty"`
}

func (o ModifyCicdConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyCicdConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyCicdConfigurationRequestBody", string(data)}, " ")
}
