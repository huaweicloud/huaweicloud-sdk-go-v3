package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEmailTemplateRequest Request Object
type DeleteEmailTemplateRequest struct {
	Body *DeleteEmailTemplateRequestBody `json:"body,omitempty"`
}

func (o DeleteEmailTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEmailTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteEmailTemplateRequest", string(data)}, " ")
}
