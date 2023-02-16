package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 调用成功时为图片标签内容。调用失败时无此字段。
type RecaptureDetectResponseResult struct {

	// 总体的结论： true：真实 false：虚假 uncertainty：不确定
	Suggestion *string `json:"suggestion,omitempty"`

	// 标签（如果suggestion为真时，则该值为空字符串，否则不为空）。recapture：翻拍图
	Category *string `json:"category,omitempty"`

	// 总体置信度，取值范围（0~1）。
	Score *string `json:"score,omitempty"`

	// 识别结果详情。
	Detail *[]RecaptureDetectResponseResultDetail `json:"detail,omitempty"`
}

func (o RecaptureDetectResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecaptureDetectResponseResult struct{}"
	}

	return strings.Join([]string{"RecaptureDetectResponseResult", string(data)}, " ")
}
