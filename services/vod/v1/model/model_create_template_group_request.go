package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplateGroupRequest Request Object
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
