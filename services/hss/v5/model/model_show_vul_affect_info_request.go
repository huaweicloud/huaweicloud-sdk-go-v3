package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulAffectInfoRequest Request Object
type ShowVulAffectInfoRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 主机ID，不传时查询该漏洞影响范围内的所有主机 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: - linux_vul：Linux漏洞。 - windows_vul：Windows漏洞。 - web_cms：Web-CMS漏洞。 - app_vul：应用漏洞。 - urgent_vul：应急漏洞。  **默认取值**: 不涉及
	VulType *string `json:"vul_type,omitempty"`

	// **参数解释**: 漏洞ID **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	VulId string `json:"vul_id"`

	// **参数解释**: 漏洞处理状态 **约束限制**: 不涉及 **取值范围**: - handled：已处理。 - unhandled：未处理。  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`
}

func (o ShowVulAffectInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulAffectInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowVulAffectInfoRequest", string(data)}, " ")
}
