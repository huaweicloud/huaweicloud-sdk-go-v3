package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunEntitySentimentResponse struct {
	// 响应的文本

	Content *string `json:"content,omitempty"`
	// 响应的实体

	Entity *string `json:"entity,omitempty"`
	// 响应的情感标签，0表示负面，1表示非负面，2表示不相关

	Label *int32 `json:"label,omitempty"`
	// 该实体在文本中的情感label的置信度

	Confidence *float64 `json:"confidence,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunEntitySentimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEntitySentimentResponse struct{}"
	}

	return strings.Join([]string{"RunEntitySentimentResponse", string(data)}, " ")
}
