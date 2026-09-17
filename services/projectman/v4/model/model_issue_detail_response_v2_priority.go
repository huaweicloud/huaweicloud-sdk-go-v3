package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Priority **参数解释：** 工作项的优先级。
type IssueDetailResponseV2Priority struct {

	// **参数解释：** 工作项的优先级id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的优先级名称。 **取值范围：** 高，中，低。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Priority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Priority struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Priority", string(data)}, " ")
}
