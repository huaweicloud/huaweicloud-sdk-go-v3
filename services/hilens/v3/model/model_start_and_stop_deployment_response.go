package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAndStopDeploymentResponse Response Object
type StartAndStopDeploymentResponse struct {

	// 部署Id
	DeploymentId   *string `json:"deployment_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAndStopDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAndStopDeploymentResponse struct{}"
	}

	return strings.Join([]string{"StartAndStopDeploymentResponse", string(data)}, " ")
}
