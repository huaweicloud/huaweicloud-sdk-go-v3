package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVmsSendTaskRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateVmsSendTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVmsSendTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsSendTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVmsSendTaskRequest", string(data)}, " ")
}
