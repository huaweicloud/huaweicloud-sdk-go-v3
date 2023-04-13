package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssetSummaryRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间.格式为(YYYYMMDD'T'HHMMSS'Z')。 取值为当前系统的GMT时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *ListAssetSummarysReq `json:"body,omitempty"`
}

func (o ListAssetSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListAssetSummaryRequest", string(data)}, " ")
}
