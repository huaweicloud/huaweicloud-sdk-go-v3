package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScrumPlanToProjectRequest Request Object
type CreateScrumPlanToProjectRequest struct {

	// project_id
	ProjectId string `json:"project_id"`

	Body *string `json:"body,omitempty"`
}

func (o CreateScrumPlanToProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScrumPlanToProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateScrumPlanToProjectRequest", string(data)}, " ")
}
