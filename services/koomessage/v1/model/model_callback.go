package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Callback 注册详情。
type Callback struct {

	// 回调地址。
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 注册回调的唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 回调类型。  - 0：智能信息发送回调 - 1：模板状态回调
	UrlType *int32 `json:"url_type,omitempty"`
}

func (o Callback) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Callback struct{}"
	}

	return strings.Join([]string{"Callback", string(data)}, " ")
}
