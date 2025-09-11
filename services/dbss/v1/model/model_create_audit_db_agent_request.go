package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuditDbAgentRequest Request Object
type CreateAuditDbAgentRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// **参数解释**： 审计Agent ID。可通过获取agent列表接口agent_id字段获取 **约束限制**： 不涉及 **取值范围**： 以获取agent列表接口agent_id值为准，字符长度16-64。 **默认取值**： 不涉及
	AgentId string `json:"agent_id"`

	Body *AgentEditRequest `json:"body,omitempty"`
}

func (o CreateAuditDbAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuditDbAgentRequest struct{}"
	}

	return strings.Join([]string{"CreateAuditDbAgentRequest", string(data)}, " ")
}
