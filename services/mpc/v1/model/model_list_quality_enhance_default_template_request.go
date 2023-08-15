package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQualityEnhanceDefaultTemplateRequest Request Object
type ListQualityEnhanceDefaultTemplateRequest struct {
}

func (o ListQualityEnhanceDefaultTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityEnhanceDefaultTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListQualityEnhanceDefaultTemplateRequest", string(data)}, " ")
}
