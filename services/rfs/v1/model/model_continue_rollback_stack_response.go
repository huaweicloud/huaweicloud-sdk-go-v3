package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueRollbackStackResponse Response Object
type ContinueRollbackStackResponse struct {

	// 继续回滚触发部署生成的唯一的deployment_id，由资源编排服务生成，通常为UUID  接受请求，进行异步处理。可以调用GetStackMetadata来获取异步请求的部署状态  **注意：** * 资源编排服务最终会将异步部署请求排队，在服务端空闲的情况下逐个处理。用户最大等待时长为6小时
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ContinueRollbackStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueRollbackStackResponse struct{}"
	}

	return strings.Join([]string{"ContinueRollbackStackResponse", string(data)}, " ")
}
