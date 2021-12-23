package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDeploymentHostResponse struct {
	// 主机id

	HostId         *string `json:"host_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeploymentHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentHostResponse", string(data)}, " ")
}
