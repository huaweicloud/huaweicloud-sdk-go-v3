package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDeploymentHostResponse Response Object
type CreateDeploymentHostResponse struct {

	// 主机id
	HostId         *string `json:"host_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeploymentHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentHostResponse struct{}"
	}

	return strings.Join([]string{"CreateDeploymentHostResponse", string(data)}, " ")
}
