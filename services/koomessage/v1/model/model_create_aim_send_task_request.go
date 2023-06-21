package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAimSendTaskRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateAimSendTaskRequestBody `json:"body,omitempty"`
}

func (o CreateAimSendTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimSendTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAimSendTaskRequest", string(data)}, " ")
}
