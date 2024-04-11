package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployStackResponse Response Object
type DeployStackResponse struct {

	// 部署ID 接受请求，进行异步处理。可以调用GetStackMetadata来获取异步请求的部署状态  **注意：** * 部署资源栈后，资源编排服务持久化请求并立即返回，客户端不等待请求最终处理完成，用户无法实时感知请求处理结果 * 资源编排服务最终会将异步部署请求排队，在服务端空闲的情况下逐个处理。用户最大等待时长为6小时
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeployStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployStackResponse struct{}"
	}

	return strings.Join([]string{"DeployStackResponse", string(data)}, " ")
}
