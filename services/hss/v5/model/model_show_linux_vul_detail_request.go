package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLinuxVulDetailRequest Request Object
type ShowLinuxVulDetailRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit int32 `json:"limit"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 不涉及
	Offset int32 `json:"offset"`

	// **参数解释**: 漏洞补丁编号 **约束限制**: 不涉及 **取值范围**: 字符长度1-256 **默认取值**: 不涉及
	VulId string `json:"vul_id"`

	// **参数解释**: 漏洞编号 **约束限制**: 不涉及 **取值范围**: 字符长度1-32 **默认取值**: 不涉及
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: 漏洞处置状态 **约束限制**: 不涉及 **取值范围**: - unhandled：未处理 - handled：已处理  **默认取值**: 不涉及
	HandleStatus *string `json:"handle_status,omitempty"`
}

func (o ShowLinuxVulDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinuxVulDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowLinuxVulDetailRequest", string(data)}, " ")
}
