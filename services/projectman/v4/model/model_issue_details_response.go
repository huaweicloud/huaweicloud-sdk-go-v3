package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailsResponse 工作项详情响应
type IssueDetailsResponse struct {

	// **参数解释**： 工作项ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项编号。 **取值范围**： 不涉及。
	Number *string `json:"number,omitempty"`

	// **参数解释**： 工作项类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 停留天数。 **取值范围**： 不涉及。
	StayDays *int32 `json:"stay_days,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 工作项创建时间。 **取值范围**： 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释**： 工作项标题。 **取值范围**： 不涉及。
	Title *string `json:"title,omitempty"`

	SecurityLevel *SecurityLevelResult `json:"security_level,omitempty"`
}

func (o IssueDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailsResponse struct{}"
	}

	return strings.Join([]string{"IssueDetailsResponse", string(data)}, " ")
}
