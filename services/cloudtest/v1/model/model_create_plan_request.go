package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePlanRequest struct {
	Body *CreatePlanRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreatePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlanRequest struct{}"
	}

	return strings.Join([]string{"CreatePlanRequest", string(data)}, " ")
}
