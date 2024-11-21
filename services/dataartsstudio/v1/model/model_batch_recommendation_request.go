package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRecommendationRequest struct {

	// 标签信息。
	Recommendations *[]Recommendation `json:"recommendations,omitempty"`

	// 资产guid。
	Guids *[]string `json:"guids,omitempty"`

	// 添加资产类型。cover：覆盖  追加：append。默认追加。
	AddType *string `json:"add_type,omitempty"`
}

func (o BatchRecommendationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRecommendationRequest struct{}"
	}

	return strings.Join([]string{"BatchRecommendationRequest", string(data)}, " ")
}
