package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV2ParentIssue **参数解释：** 工作项的父工作项信息。
type IssueDetailResponseV2ParentIssue struct {

	// **参数解释：** 工作项的父工作项id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的父工作项名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV2ParentIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2ParentIssue struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2ParentIssue", string(data)}, " ")
}
