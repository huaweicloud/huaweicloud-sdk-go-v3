package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStructTemplateRequest struct {
	Body *DeleteStructTemplateReqBody `json:"body,omitempty"`
}

func (o DeleteStructTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStructTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteStructTemplateRequest", string(data)}, " ")
}
