package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResolveTaskRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateResolveTaskRequestBody `json:"body,omitempty"`
}

func (o CreateResolveTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolveTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateResolveTaskRequest", string(data)}, " ")
}
