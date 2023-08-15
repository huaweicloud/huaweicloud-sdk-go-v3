package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextDetectionDataReq struct {

	// 待检测文本，编码格式为“utf-8”，限定2000个字符以内，文本长度超过1500个字符时，只检测前1500个字符。
	Text string `json:"text"`
}

func (o TextDetectionDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionDataReq struct{}"
	}

	return strings.Join([]string{"TextDetectionDataReq", string(data)}, " ")
}
