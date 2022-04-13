package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStructTemplateRequest struct {
	Body *LtsStructTemplateInfo `json:"body,omitempty"`
}

func (o CreateStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateStructTemplateRequest", string(data)}, " ")
}
