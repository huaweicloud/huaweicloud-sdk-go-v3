package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetSummaryRequest Request Object
type ListAssetSummaryRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *ListAssetSummarysReq `json:"body,omitempty"`
}

func (o ListAssetSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListAssetSummaryRequest", string(data)}, " ")
}
