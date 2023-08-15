package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeploymentHostResponse Response Object
type UpdateDeploymentHostResponse struct {

	// 主机id
	HostId         *string `json:"host_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeploymentHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentHostResponse", string(data)}, " ")
}
