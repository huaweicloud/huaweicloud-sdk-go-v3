package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletesResponseResultDeleteIssue **参数解释：** 删除的工作项。 **取值范围：** 不涉及。
type BatchDeletesResponseResultDeleteIssue struct {

	// **参数解释：** 删除的工作项id。 **取值范围：** 不涉及。
	DelIssueId *[]int32 `json:"del_issue_id,omitempty"`

	// **参数解释：** 删除的工作项详情。 **取值范围：** 不涉及。
	DelIssue *[]BatchDeletesResponseResultDeleteIssueDelIssue `json:"del_issue,omitempty"`
}

func (o BatchDeletesResponseResultDeleteIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletesResponseResultDeleteIssue struct{}"
	}

	return strings.Join([]string{"BatchDeletesResponseResultDeleteIssue", string(data)}, " ")
}
