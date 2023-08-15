package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobTemplatesRequest Request Object
type CreateJobTemplatesRequest struct {
	Body *CreateJobTemplatesRequestBody `json:"body,omitempty"`
}

func (o CreateJobTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobTemplatesRequest struct{}"
	}

	return strings.Join([]string{"CreateJobTemplatesRequest", string(data)}, " ")
}
