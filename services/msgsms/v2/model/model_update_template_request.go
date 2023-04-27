package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTemplateRequest struct {

	// 模板ID
	Id string `json:"id"`

	Body *SmsTemplateReq `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}
