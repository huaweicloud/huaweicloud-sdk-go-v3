package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupplyRecommendationRequest Request Object
type ListSupplyRecommendationRequest struct {
	Body *ListSupplyRecommendationRequestBody `json:"body,omitempty"`
}

func (o ListSupplyRecommendationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplyRecommendationRequest struct{}"
	}

	return strings.Join([]string{"ListSupplyRecommendationRequest", string(data)}, " ")
}
