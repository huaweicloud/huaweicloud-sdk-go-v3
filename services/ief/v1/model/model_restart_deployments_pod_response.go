package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartDeploymentsPodResponse Response Object
type RestartDeploymentsPodResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartDeploymentsPodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartDeploymentsPodResponse struct{}"
	}

	return strings.Join([]string{"RestartDeploymentsPodResponse", string(data)}, " ")
}
