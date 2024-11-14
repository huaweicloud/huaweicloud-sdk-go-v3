package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartBeautyPreviewJobResponse Response Object
type StartBeautyPreviewJobResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartBeautyPreviewJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartBeautyPreviewJobResponse struct{}"
	}

	return strings.Join([]string{"StartBeautyPreviewJobResponse", string(data)}, " ")
}
