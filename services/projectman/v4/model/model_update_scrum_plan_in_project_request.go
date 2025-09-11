package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumPlanInProjectRequest Request Object
type UpdateScrumPlanInProjectRequest struct {

	// project_id
	ProjectId string `json:"project_id"`

	// plan_id
	PlanId string `json:"plan_id"`

	Body *string `json:"body,omitempty"`
}

func (o UpdateScrumPlanInProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumPlanInProjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateScrumPlanInProjectRequest", string(data)}, " ")
}
