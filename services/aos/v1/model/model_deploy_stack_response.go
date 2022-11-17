package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeployStackResponse struct {

	// 部署ID
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
