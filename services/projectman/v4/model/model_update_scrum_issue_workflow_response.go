package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumIssueWorkflowResponse Response Object
type UpdateScrumIssueWorkflowResponse struct {
	Result *UpdateIssueFlowsResponseResult `json:"result,omitempty"`

	// **参数解释：** 工作项状态流转的返回状态。 **取值范围：** success：返回成功。 error：返回失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScrumIssueWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumIssueWorkflowResponse struct{}"
	}

	return strings.Join([]string{"UpdateScrumIssueWorkflowResponse", string(data)}, " ")
}
