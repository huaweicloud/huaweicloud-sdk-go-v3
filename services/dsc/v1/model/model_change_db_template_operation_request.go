package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeDbTemplateOperationRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`

	Body *MaskSwitchRequest `json:"body,omitempty"`
}

func (o ChangeDbTemplateOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDbTemplateOperationRequest struct{}"
	}

	return strings.Join([]string{"ChangeDbTemplateOperationRequest", string(data)}, " ")
}
