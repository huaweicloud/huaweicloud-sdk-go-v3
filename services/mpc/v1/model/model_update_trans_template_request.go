package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTransTemplateRequest struct {
	Body *ModifyTransTemplateReq `json:"body,omitempty"`
}

func (o UpdateTransTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateRequest", string(data)}, " ")
}
