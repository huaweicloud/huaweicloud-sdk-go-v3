package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageWhiteListsRequest Request Object
type ListImageWhiteListsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像。 - registry：仓库镜像。  **默认取值**: 不涉及
	GlobalImageType string `json:"global_image_type"`

	// **参数解释**： 白名单类型 **约束限制**: 不涉及 **取值范围**: - vulnerability：漏洞白名单。  **默认取值**: 不涉及
	Type string `json:"type"`

	// **参数解释**： 漏洞名称（只在查询漏洞白名单时使用） **约束限制**： 不涉及 **取值范围**： 字符长度0-256位 **默认取值**： 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**： 漏洞类型（只在查询漏洞白名单时使用） **约束限制**: 不涉及 **取值范围**: - linux_vul：linux漏洞。 - app_vul：应用漏洞。  **默认取值**: 不涉及
	VulType *string `json:"vul_type,omitempty"`
}

func (o ListImageWhiteListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageWhiteListsRequest struct{}"
	}

	return strings.Join([]string{"ListImageWhiteListsRequest", string(data)}, " ")
}
