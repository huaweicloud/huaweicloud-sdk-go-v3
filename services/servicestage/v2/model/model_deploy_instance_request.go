package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployInstanceRequest Request Object
type DeployInstanceRequest struct {
	Body *InstanceDeployment `json:"body,omitempty"`
}

func (o DeployInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeployInstanceRequest", string(data)}, " ")
}
