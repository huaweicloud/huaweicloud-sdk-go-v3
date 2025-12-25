package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIndicatorDetailResponse Response Object
type ShowIndicatorDetailResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *IndicatorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIndicatorDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndicatorDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowIndicatorDetailResponse", string(data)}, " ")
}
