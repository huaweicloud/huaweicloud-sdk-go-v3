package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunFileTranslationResponse struct {

	// 创建的任务标识, 如果创建任务成功时必存在。调用失败时无此字段。
	JobId *string `json:"job_id,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunFileTranslationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFileTranslationResponse struct{}"
	}

	return strings.Join([]string{"RunFileTranslationResponse", string(data)}, " ")
}
