package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQualityEnhanceTemplateRequest Request Object
type UpdateQualityEnhanceTemplateRequest struct {
	Body *UpdateQualityEnhanceTemplateReq `json:"body,omitempty"`
}

func (o UpdateQualityEnhanceTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateRequest", string(data)}, " ")
}
