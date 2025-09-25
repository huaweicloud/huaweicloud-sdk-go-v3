package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageWhiteListsRequest Request Object
type DeleteImageWhiteListsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**: - local：本地镜像。 - registry：仓库镜像。  **默认取值**: 不涉及
	GlobalImageType string `json:"global_image_type"`

	// **参数解释**： 白名单类型 **约束限制**: 不涉及 **取值范围**: - vulnerability：漏洞白名单。  **默认取值**: 不涉及
	Type string `json:"type"`

	Body *DeleteImageWhiteListsRequestBody `json:"body,omitempty"`
}

func (o DeleteImageWhiteListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageWhiteListsRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageWhiteListsRequest", string(data)}, " ")
}
