package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueListResult 工作项列表查询结果
type IssueListResult struct {

	// **参数解释：** 工作项列表。 **取值范围：** 不涉及。
	Issues *[]IssueVo `json:"issues,omitempty"`

	// **参数解释：** 符合过滤条件的工作项总数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o IssueListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueListResult struct{}"
	}

	return strings.Join([]string{"IssueListResult", string(data)}, " ")
}
