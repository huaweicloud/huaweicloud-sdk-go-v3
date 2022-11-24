package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWorkflowExecutionForPageResponse struct {
	Pager *Pager `json:"pager,omitempty"`

	HisRecords     *FlowExecutionBriefV2 `json:"hisRecords,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowWorkflowExecutionForPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionForPageResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionForPageResponse", string(data)}, " ")
}
