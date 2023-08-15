package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunEventExtractionResponse Response Object
type RunEventExtractionResponse struct {

	// 事件抽取结果。调用失败时无此字段。
	Events *[]EventExtractionResponseItem `json:"events,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunEventExtractionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEventExtractionResponse struct{}"
	}

	return strings.Join([]string{"RunEventExtractionResponse", string(data)}, " ")
}
