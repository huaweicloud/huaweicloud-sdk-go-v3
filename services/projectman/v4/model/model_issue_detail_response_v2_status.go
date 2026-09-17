package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Status **参数解释：** 工作项的当前状态。
type IssueDetailResponseV2Status struct {

	// **参数解释：** 工作项的当前状态id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的当前状态名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Status struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Status", string(data)}, " ")
}
