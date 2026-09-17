package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIssueFlowsResponseResult **参数解释：** 工作项状态流转的返回结果。
type UpdateIssueFlowsResponseResult struct {
	Issue *IssueNew `json:"issue,omitempty"`
}

func (o UpdateIssueFlowsResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIssueFlowsResponseResult struct{}"
	}

	return strings.Join([]string{"UpdateIssueFlowsResponseResult", string(data)}, " ")
}
