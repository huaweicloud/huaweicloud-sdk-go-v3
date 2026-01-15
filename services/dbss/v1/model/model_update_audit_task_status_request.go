package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditTaskStatusRequest Request Object
type UpdateAuditTaskStatusRequest struct {

	// **参数解释**： 业务类型。 **约束限制**： 区分大小写，以取值范围为准 **取值范围**： - risk：风险  - audit：审计实例 **默认取值**： 不涉及
	BusiType string `json:"busi_type"`

	Body *AuditSummaryStatusRequest `json:"body,omitempty"`
}

func (o UpdateAuditTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuditTaskStatusRequest", string(data)}, " ")
}
