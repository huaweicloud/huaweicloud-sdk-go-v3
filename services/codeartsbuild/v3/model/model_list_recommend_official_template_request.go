package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecommendOfficialTemplateRequest Request Object
type ListRecommendOfficialTemplateRequest struct {
	Body *ListRecommendOfficialTemplateRequestBody `json:"body,omitempty"`
}

func (o ListRecommendOfficialTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecommendOfficialTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListRecommendOfficialTemplateRequest", string(data)}, " ")
}
