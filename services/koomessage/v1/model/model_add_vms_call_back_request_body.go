package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddVmsCallBackRequestBody struct {

	// 回调类型。  - 0：发送状态回执 - 1：上行消息回执
	UrlType int32 `json:"url_type"`

	// 回调地址，必须是http或https开头的字符串，不能为空。  > 建议使用https。
	CallbackUrl string `json:"callback_url"`

	// 备注。
	Remark *string `json:"remark,omitempty"`
}

func (o AddVmsCallBackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVmsCallBackRequestBody struct{}"
	}

	return strings.Join([]string{"AddVmsCallBackRequestBody", string(data)}, " ")
}
