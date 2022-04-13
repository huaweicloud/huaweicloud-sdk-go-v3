package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunClassificationResponse struct {
	Result *ClassificationResult `json:"result,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunClassificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunClassificationResponse struct{}"
	}

	return strings.Join([]string{"RunClassificationResponse", string(data)}, " ")
}
