package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VmsCallBack 回执接口。
type VmsCallBack struct {

	// 回调类型。  - 0：发送状态回执 - 1：上行消息回执
	UrlType *int32 `json:"url_type,omitempty"`

	// 回调地址。
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o VmsCallBack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VmsCallBack struct{}"
	}

	return strings.Join([]string{"VmsCallBack", string(data)}, " ")
}
