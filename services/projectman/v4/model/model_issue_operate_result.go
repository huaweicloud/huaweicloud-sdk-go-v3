package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueOperateResult 工作项操作返回值
type IssueOperateResult struct {

	// **参数解释**： 变更的工作项ID。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项变更人ID。 **取值范围**： 不涉及
	Operator *string `json:"operator,omitempty"`

	// **参数解释**： 工作项的作废标识，枚举类型。 **取值范围**： - 正在工作：可正常操作的工作项 - 作废：软删除后的工作项，可在回收站恢复 - 删除：彻底删除后的工作项，无法恢复
	State *string `json:"state,omitempty"`

	// **参数解释**： 工作项变更时间。 **取值范围**： 不涉及
	OperateTime *string `json:"operate_time,omitempty"`
}

func (o IssueOperateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueOperateResult struct{}"
	}

	return strings.Join([]string{"IssueOperateResult", string(data)}, " ")
}
