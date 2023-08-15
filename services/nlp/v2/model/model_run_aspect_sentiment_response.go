package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAspectSentimentResponse Response Object
type RunAspectSentimentResponse struct {

	// 待分析文本
	Text *string `json:"text,omitempty"`

	// 该文本的整体情感标签，取值如下： 0  负向 1  正向
	Label *int32 `json:"label,omitempty"`

	// 该文本整体情感label的置信度,小数点精确到3位。
	Confidence *float32 `json:"confidence,omitempty"`

	// 属性级情感挖掘列表
	AspectOpinions *[]AspectOpinion `json:"aspect_opinions,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunAspectSentimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAspectSentimentResponse struct{}"
	}

	return strings.Join([]string{"RunAspectSentimentResponse", string(data)}, " ")
}
