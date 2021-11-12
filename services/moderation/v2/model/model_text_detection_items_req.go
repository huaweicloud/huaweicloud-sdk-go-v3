package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TextDetectionItemsReq struct {
	// 待检测文本，编码格式为“utf-8”，限定5000个字符以内，文本长度超过5000个字符时，只检测前5000个字符。

	Text string `json:"text"`
	// 文本类型，默认为“content”，即正文内容，当前只支持“content”类型，未来会扩大支持类型范围。

	Type *string `json:"type,omitempty"`
}

func (o TextDetectionItemsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionItemsReq struct{}"
	}

	return strings.Join([]string{"TextDetectionItemsReq", string(data)}, " ")
}
