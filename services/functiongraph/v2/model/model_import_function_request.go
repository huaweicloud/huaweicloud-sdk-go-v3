package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFunctionRequest Request Object
type ImportFunctionRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *ImportFunctionRequestBody `json:"body,omitempty"`
}

func (o ImportFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ImportFunctionRequest", string(data)}, " ")
}
