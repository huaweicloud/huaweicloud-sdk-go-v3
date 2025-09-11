package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditTaskStatusRequest Request Object
type ShowAuditTaskStatusRequest struct {

	// **参数解释**： 业务类型。 **约束限制**： 不涉及 **取值范围**：   - audit：审计实例    - risk：风险 **默认取值**： 不涉及
	BusiType string `json:"busi_type"`
}

func (o ShowAuditTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditTaskStatusRequest", string(data)}, " ")
}
