package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2Iteration **参数解释：** 工作项所属迭代信息。
type IssueDetailResponseV2Iteration struct {

	// **参数解释：** 工作项的迭代id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的迭代名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2Iteration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2Iteration struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2Iteration", string(data)}, " ")
}
