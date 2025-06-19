package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecommendOfficialTemplateResponse Response Object
type ListRecommendOfficialTemplateResponse struct {
	Result *QueryTemplatesResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecommendOfficialTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecommendOfficialTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListRecommendOfficialTemplateResponse", string(data)}, " ")
}
