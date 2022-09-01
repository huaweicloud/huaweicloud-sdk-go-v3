package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunAspectSentimentAdvanceResponse struct {

	// 待分析文本
	Content *string `json:"content,omitempty" xml:"content"`

	// 该文本的整体情感标签，取值如下： 0  负向 1  正向
	Label *int32 `json:"label,omitempty" xml:"label"`

	// 该文本整体情感label的置信度,小数点精确到3位。
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence"`

	// 属性情感挖掘列表
	AspectOpinions *[]AspectAdvanceOpinion `json:"aspect_opinions,omitempty" xml:"aspect_opinions"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o RunAspectSentimentAdvanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAspectSentimentAdvanceResponse struct{}"
	}

	return strings.Join([]string{"RunAspectSentimentAdvanceResponse", string(data)}, " ")
}
