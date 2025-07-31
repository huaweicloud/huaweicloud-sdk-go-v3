package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTtscGroupAssetsRequest Request Object
type SetTtscGroupAssetsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 分组id
	GroupId string `json:"group_id"`

	Body *SetTtscGroupAssetsRequestBody `json:"body,omitempty"`
}

func (o SetTtscGroupAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTtscGroupAssetsRequest struct{}"
	}

	return strings.Join([]string{"SetTtscGroupAssetsRequest", string(data)}, " ")
}
