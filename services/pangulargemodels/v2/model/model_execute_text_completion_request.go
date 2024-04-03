package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTextCompletionRequest Request Object
type ExecuteTextCompletionRequest struct {

	// 发送的实体的MIME类型。
	ContentType *string `json:"Content-Type,omitempty"`

	// 资源池ID
	PoolId string `json:"pool_id"`

	// 模型的部署ID
	DeploymentId string `json:"deployment_id"`

	Body *TextCompletionReq `json:"body,omitempty"`
}

func (o ExecuteTextCompletionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTextCompletionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTextCompletionRequest", string(data)}, " ")
}
