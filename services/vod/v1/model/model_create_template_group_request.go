package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTemplateGroupRequest struct {
	Body *TransTemplateGroup `json:"body,omitempty"`
}

func (o CreateTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupRequest", string(data)}, " ")
}
