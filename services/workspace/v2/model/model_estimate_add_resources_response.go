package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateAddResourcesResponse Response Object
type EstimateAddResourcesResponse struct {

	// 币种，比如CNY
	Currency *string `json:"currency,omitempty"`

	// 询价结果
	CloudServiceRatingResults *[]CloudServiceRatingResult `json:"cloud_service_rating_results,omitempty"`
	HttpStatusCode            int                         `json:"-"`
}

func (o EstimateAddResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateAddResourcesResponse struct{}"
	}

	return strings.Join([]string{"EstimateAddResourcesResponse", string(data)}, " ")
}
