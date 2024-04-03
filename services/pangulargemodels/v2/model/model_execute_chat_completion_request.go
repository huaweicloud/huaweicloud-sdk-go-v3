package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteChatCompletionRequest Request Object
type ExecuteChatCompletionRequest struct {

	// 发送的实体的MIME类型。
	ContentType *string `json:"Content-Type,omitempty"`

	// 资源池ID
	PoolId string `json:"pool_id"`

	// 模型的部署ID
	DeploymentId string `json:"deployment_id"`

	Body *ChatCompletionReq `json:"body,omitempty"`
}

func (o ExecuteChatCompletionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteChatCompletionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteChatCompletionRequest", string(data)}, " ")
}
