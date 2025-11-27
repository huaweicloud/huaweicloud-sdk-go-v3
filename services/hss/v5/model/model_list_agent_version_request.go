package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentVersionRequest Request Object
type ListAgentVersionRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 操作系统类型。 **约束限制**: 不涉及 **取值范围**: - Linux：Linux操作系统。 - Windows：Windows操作系统。  **默认取值**: 无
	OsType *string `json:"os_type,omitempty"`
}

func (o ListAgentVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentVersionRequest struct{}"
	}

	return strings.Join([]string{"ListAgentVersionRequest", string(data)}, " ")
}
