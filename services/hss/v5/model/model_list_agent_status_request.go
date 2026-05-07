package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentStatusRequest Request Object
type ListAgentStatusRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: Agent的唯一标识ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId string `json:"agent_id"`

	// **参数解释**： agent状态 **约束限制**: 不涉及 **取值范围**: -not_installed：未安装 -online：在线 -offline：离线 -install_failed：安装失败 -installing：安装中  **默认取值**: 不涉及
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 异常原因 **约束限制**: 不涉及 **取值范围**: 字符长度0-512位 **默认取值**: 不涉及
	AbnormalReason *string `json:"abnormal_reason,omitempty"`
}

func (o ListAgentStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAgentStatusRequest", string(data)}, " ")
}
