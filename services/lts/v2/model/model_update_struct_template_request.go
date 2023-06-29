package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStructTemplateRequest Request Object
type UpdateStructTemplateRequest struct {
	Body *LtsStructTemplateInfo `json:"body,omitempty"`
}

func (o UpdateStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateStructTemplateRequest", string(data)}, " ")
}
