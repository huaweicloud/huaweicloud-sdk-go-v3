package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviewByPageResponse Response Object
type ShowReviewByPageResponse struct {
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoReview `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReviewByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviewByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowReviewByPageResponse", string(data)}, " ")
}
