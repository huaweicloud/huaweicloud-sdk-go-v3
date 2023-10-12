package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactorySupplementDataInstanceRequest Request Object
type CreateFactorySupplementDataInstanceRequest struct {

	// 当前作业的空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 定义消息体类型的请求头
	ContentType string `json:"Content-Type"`

	Body *CreateFactorySupplementDataInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateFactorySupplementDataInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceRequest", string(data)}, " ")
}
