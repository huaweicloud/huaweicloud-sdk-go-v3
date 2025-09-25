package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogstashTemplateRequest Request Object
type DeleteLogstashTemplateRequest struct {
	Body *DeleteTemplateReqNew `json:"body,omitempty"`
}

func (o DeleteLogstashTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogstashTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogstashTemplateRequest", string(data)}, " ")
}
