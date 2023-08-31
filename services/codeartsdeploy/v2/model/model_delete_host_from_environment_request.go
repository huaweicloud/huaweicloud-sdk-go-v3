package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostFromEnvironmentRequest Request Object
type DeleteHostFromEnvironmentRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o DeleteHostFromEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostFromEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostFromEnvironmentRequest", string(data)}, " ")
}
