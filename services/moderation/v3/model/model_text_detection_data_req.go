package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextDetectionDataReq struct {

	// 待检测文本，编码格式为“utf-8”，限定2000个字符以内，文本长度超过2000个字符时，只检测前2000个字符。
	Text string `json:"text" xml:"text"`
}

func (o TextDetectionDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionDataReq struct{}"
	}

	return strings.Join([]string{"TextDetectionDataReq", string(data)}, " ")
}
