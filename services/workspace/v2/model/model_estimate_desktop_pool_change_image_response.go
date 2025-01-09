package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolChangeImageResponse Response Object
type EstimateDesktopPoolChangeImageResponse struct {

	// 币种，比如CNY
	Currency *string `json:"currency,omitempty"`

	// 询价结果
	CloudServiceRatingResults *[]CloudServiceRatingResult `json:"cloud_service_rating_results,omitempty"`
	HttpStatusCode            int                         `json:"-"`
}

func (o EstimateDesktopPoolChangeImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolChangeImageResponse struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolChangeImageResponse", string(data)}, " ")
}
