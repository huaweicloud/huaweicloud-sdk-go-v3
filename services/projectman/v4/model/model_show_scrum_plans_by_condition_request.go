package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScrumPlansByConditionRequest Request Object
type ShowScrumPlansByConditionRequest struct {

	// project_id
	ProjectId string `json:"project_id"`

	Body *string `json:"body,omitempty"`
}

func (o ShowScrumPlansByConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScrumPlansByConditionRequest struct{}"
	}

	return strings.Join([]string{"ShowScrumPlansByConditionRequest", string(data)}, " ")
}
