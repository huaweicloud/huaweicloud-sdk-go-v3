package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTemplateGroupRequest struct {
	Body *ModifyTransTemplateGroup `json:"body,omitempty"`
}

func (o UpdateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupRequest", string(data)}, " ")
}
