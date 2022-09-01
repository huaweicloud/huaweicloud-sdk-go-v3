package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type PostEventExtractionReq struct {

	// 待分析文本，长度为1~256，文本编码为UTF-8。
	Text string `json:"text" xml:"text"`
}

func (o PostEventExtractionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostEventExtractionReq struct{}"
	}

	return strings.Join([]string{"PostEventExtractionReq", string(data)}, " ")
}
