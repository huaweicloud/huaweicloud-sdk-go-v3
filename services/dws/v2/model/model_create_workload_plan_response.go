package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateWorkloadPlanResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWorkloadPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkloadPlanResponse", string(data)}, " ")
}
