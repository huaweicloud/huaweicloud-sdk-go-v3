package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCallbackRequestBody 回调地址请求体。
type AddCallbackRequestBody struct {

	// 回调地址。  > - 必须http或https开头，建议使用https > - 支持域名或公网IP回调地址，不支持私网IP回调地址
	CallbackUrl string `json:"callback_url"`

	// 回调类型。  - 0：智能信息发送回调 - 1：模板状态回调
	UrlType *int32 `json:"url_type,omitempty"`
}

func (o AddCallbackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCallbackRequestBody struct{}"
	}

	return strings.Join([]string{"AddCallbackRequestBody", string(data)}, " ")
}
