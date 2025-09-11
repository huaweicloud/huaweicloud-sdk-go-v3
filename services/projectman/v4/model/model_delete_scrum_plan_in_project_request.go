package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScrumPlanInProjectRequest Request Object
type DeleteScrumPlanInProjectRequest struct {

	// project_id
	ProjectId string `json:"project_id"`
}

func (o DeleteScrumPlanInProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScrumPlanInProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteScrumPlanInProjectRequest", string(data)}, " ")
}
