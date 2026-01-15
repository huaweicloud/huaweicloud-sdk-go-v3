package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditTagsRequest Request Object
type ListAuditTagsRequest struct {

	// **参数解释**：  资源类型。 **约束限制**： 不涉及 **取值范围**：  - auditInstance：审计 **默认取值**： 不涉及
	ResourceType string `json:"resource_type"`

	// **参数解释**：  资源ID。可在查询实例列表接口的resource_id字段获取。 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口获取值为准，字符长度32-64。 **默认取值**： 不涉及
	ResourceId string `json:"resource_id"`
}

func (o ListAuditTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditTagsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditTagsRequest", string(data)}, " ")
}
