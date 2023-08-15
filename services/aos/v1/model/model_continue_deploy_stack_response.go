package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueDeployStackResponse Response Object
type ContinueDeployStackResponse struct {

	// 部署ID
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ContinueDeployStackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueDeployStackResponse struct{}"
	}

	return strings.Join([]string{"ContinueDeployStackResponse", string(data)}, " ")
}
