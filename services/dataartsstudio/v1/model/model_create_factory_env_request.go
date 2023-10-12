package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactoryEnvRequest Request Object
type CreateFactoryEnvRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 消息体类型的请求头
	ContentType string `json:"Content-Type"`

	Body *EnvRequestBody `json:"body,omitempty"`
}

func (o CreateFactoryEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactoryEnvRequest struct{}"
	}

	return strings.Join([]string{"CreateFactoryEnvRequest", string(data)}, " ")
}
