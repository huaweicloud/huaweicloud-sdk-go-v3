package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetRequest Request Object
type DeleteAssetRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 资产ID。
	AssetId string `json:"asset_id"`

	// 删除模式
	Mode *string `json:"mode,omitempty"`
}

func (o DeleteAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetRequest", string(data)}, " ")
}
