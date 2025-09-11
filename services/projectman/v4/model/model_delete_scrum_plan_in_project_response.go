package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScrumPlanInProjectResponse Response Object
type DeleteScrumPlanInProjectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteScrumPlanInProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScrumPlanInProjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteScrumPlanInProjectResponse", string(data)}, " ")
}
