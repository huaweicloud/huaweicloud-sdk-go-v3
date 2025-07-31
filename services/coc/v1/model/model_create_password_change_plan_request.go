package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePasswordChangePlanRequest Request Object
type CreatePasswordChangePlanRequest struct {
	Body *CreatePasswordChangePlanRequestBody `json:"body,omitempty"`
}

func (o CreatePasswordChangePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordChangePlanRequest struct{}"
	}

	return strings.Join([]string{"CreatePasswordChangePlanRequest", string(data)}, " ")
}
