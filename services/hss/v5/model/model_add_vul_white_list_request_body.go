package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddVulWhiteListRequestBody struct {

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞 - windows_vul：windows漏洞 - web_cms：Web-CMS漏洞 - app_vul：应用漏洞  **默认取值**: 不涉及
	VulType string `json:"vul_type"`

	// **参数解释**: 漏洞ID列表 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值500 **默认取值**: 不涉及
	VulIds []string `json:"vul_ids"`

	// **参数解释**: 白名单规则类型 **约束限制**: 不涉及 **取值范围**: - all_host：白名单应用于全部主机 - specific_host：白名单应用于指定主机  **默认取值**: 不涉及
	RuleType string `json:"rule_type"`

	// **参数解释**: 主机ID列表，当rule_type为specific_host时，该字段必填 **约束限制**: 不涉及 **取值范围**: 最小值1，最大值2000 **默认取值**: 不涉及
	HostIds *[]string `json:"host_ids,omitempty"`

	// **参数解释**: 是否忽略当前已扫描出的漏洞 **约束限制**: 不涉及 **取值范围**: - true：忽略 - false：不忽略 **默认取值**: true
	WithIgnore *bool `json:"with_ignore,omitempty"`

	// **参数解释**: 白名单的描述信息 **约束限制**: 不涉及 **取值范围**: 字符长度0-1000 **默认取值**: 不涉及
	Description *string `json:"description,omitempty"`
}

func (o AddVulWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVulWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"AddVulWhiteListRequestBody", string(data)}, " ")
}
