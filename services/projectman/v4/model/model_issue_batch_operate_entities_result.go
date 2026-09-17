package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueBatchOperateEntitiesResult 工作项批量变更操作结果
type IssueBatchOperateEntitiesResult struct {

	// **参数解释**： 成功的结果。
	Success *[]IssueOperateResult `json:"success,omitempty"`

	// **参数解释**： 失败的结果。
	Failed *[]IssueOperateResult `json:"failed,omitempty"`

	// **参数解释**： 删除失败的工作项。 **取值范围**： 不涉及
	UndeletedTrees *[]IssueOperateResult `json:"undeleted_trees,omitempty"`
}

func (o IssueBatchOperateEntitiesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueBatchOperateEntitiesResult struct{}"
	}

	return strings.Join([]string{"IssueBatchOperateEntitiesResult", string(data)}, " ")
}
