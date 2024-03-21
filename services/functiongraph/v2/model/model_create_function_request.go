package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFunctionRequest Request Object
type CreateFunctionRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *CreateFunctionRequestBody `json:"body,omitempty"`
}

func (o CreateFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionRequest struct{}"
	}

	return strings.Join([]string{"CreateFunctionRequest", string(data)}, " ")
}
