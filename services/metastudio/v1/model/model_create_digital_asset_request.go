package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDigitalAssetRequest Request Object
type CreateDigitalAssetRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *CreateDigitalAssetRequestBody `json:"body,omitempty"`
}

func (o CreateDigitalAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDigitalAssetRequest struct{}"
	}

	return strings.Join([]string{"CreateDigitalAssetRequest", string(data)}, " ")
}
