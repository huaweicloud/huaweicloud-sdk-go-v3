package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Tracker **参数解释：** 工作项类型。
type IssueDetailResponseV2Tracker struct {

	// **参数解释：** 工作项的类型id。 **取值范围：** 2，3，5，6，7。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的类型名称。 **取值范围：** 2：任务/Task 3：缺陷/Bug 5：Epic 6：Feature 7：Story。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Tracker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Tracker struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Tracker", string(data)}, " ")
}
