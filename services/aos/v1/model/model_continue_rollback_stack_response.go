package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueRollbackStackResponse Response Object
type ContinueRollbackStackResponse struct {

	// 继续回滚触发部署生成的唯一的deployment_id，由资源编排服务生成，通常为UUID
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
