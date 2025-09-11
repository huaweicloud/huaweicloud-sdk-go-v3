package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumPlanInProjectResponse Response Object
type UpdateScrumPlanInProjectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScrumPlanInProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumPlanInProjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateScrumPlanInProjectResponse", string(data)}, " ")
}
