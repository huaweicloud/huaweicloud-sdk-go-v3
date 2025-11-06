package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviewSettingResponse Response Object
type ShowReviewSettingResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *ReviewSettingDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReviewSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviewSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowReviewSettingResponse", string(data)}, " ")
}
