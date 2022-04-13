package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunSummaryDomainResponse struct {
	// 根据文本请求体，返回摘要结果。调用失败时无此字段。

	Summary *string `json:"summary,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunSummaryDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSummaryDomainResponse struct{}"
	}

	return strings.Join([]string{"RunSummaryDomainResponse", string(data)}, " ")
}
