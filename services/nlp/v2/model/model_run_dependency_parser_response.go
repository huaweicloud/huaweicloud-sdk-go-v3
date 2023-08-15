package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDependencyParserResponse Response Object
type RunDependencyParserResponse struct {

	// 依存句法分析结果，词汇集合。调用失败时无此字段。
	Words *[]DependencyParserWord `json:"words,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDependencyParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDependencyParserResponse struct{}"
	}

	return strings.Join([]string{"RunDependencyParserResponse", string(data)}, " ")
}
