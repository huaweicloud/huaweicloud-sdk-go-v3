package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVoiceTrainingQuotasRequest Request Object
type ShowVoiceTrainingQuotasRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowVoiceTrainingQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVoiceTrainingQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowVoiceTrainingQuotasRequest", string(data)}, " ")
}
