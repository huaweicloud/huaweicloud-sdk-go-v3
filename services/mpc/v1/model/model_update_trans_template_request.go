package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransTemplateRequest Request Object
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
