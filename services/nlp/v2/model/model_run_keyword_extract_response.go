package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunKeywordExtractResponse struct {
	// 关键词列表。调用失败时无此字段。

	Words *[]string `json:"words,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunKeywordExtractResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunKeywordExtractResponse struct{}"
	}

	return strings.Join([]string{"RunKeywordExtractResponse", string(data)}, " ")
}
