package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCicdConfigurationRequestBody 新增cicd配置
type AddCicdConfigurationRequestBody struct {

	// **参数解释** cicd配置信息 **约束限制** 不涉及 **取值范围** 字符长度1-128位 **默认取值** 不涉及
	CicdName string `json:"cicd_name"`

	// **参数解释** 漏洞白名单 **约束限制** 不涉及 **取值范围** 最小值0，最大值1000 **默认取值** 不涉及
	VulnerabilityWhitelist *[]string `json:"vulnerability_whitelist,omitempty"`

	// **参数解释** 漏洞黑名单 **约束限制** 不涉及 **取值范围** 最小值0，最大值1000 **默认取值** 不涉及
	VulnerabilityBlocklist *[]string `json:"vulnerability_blocklist,omitempty"`

	// **参数解释** 镜像白名单 **约束限制** 不涉及 **取值范围** 最小值0，最大值1000 **默认取值** 不涉及
	ImageWhitelist *[]string `json:"image_whitelist,omitempty"`
}

func (o AddCicdConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCicdConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"AddCicdConfigurationRequestBody", string(data)}, " ")
}
