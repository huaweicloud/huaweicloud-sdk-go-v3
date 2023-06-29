package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateRequest Request Object
type UpdateTemplateRequest struct {

	// id
	Id string `json:"id"`

	// workspace 信息
	Workspace string `json:"workspace"`

	Body *TemplateRo `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}
