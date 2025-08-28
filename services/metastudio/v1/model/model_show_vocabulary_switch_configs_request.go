package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVocabularySwitchConfigsRequest Request Object
type ShowVocabularySwitchConfigsRequest struct {

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

	// 查询偏移量,若超过最大数量，则返回最后一页
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量
	Limit *int32 `json:"limit,omitempty"`

	// 类型 租户级自定义读法全局开关:vocabulary
	Key *string `json:"key,omitempty"`
}

func (o ShowVocabularySwitchConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabularySwitchConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowVocabularySwitchConfigsRequest", string(data)}, " ")
}
