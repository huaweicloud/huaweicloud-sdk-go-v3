package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupplyRecommendationResponse Response Object
type ListSupplyRecommendationResponse struct {

	// 资源供给推荐结果
	SupplyRecommendations *[]SupplyRecommendation `json:"supply_recommendations,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ListSupplyRecommendationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplyRecommendationResponse struct{}"
	}

	return strings.Join([]string{"ListSupplyRecommendationResponse", string(data)}, " ")
}
