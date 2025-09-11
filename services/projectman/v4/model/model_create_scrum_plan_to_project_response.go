package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScrumPlanToProjectResponse Response Object
type CreateScrumPlanToProjectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScrumPlanToProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScrumPlanToProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateScrumPlanToProjectResponse", string(data)}, " ")
}
