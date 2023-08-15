package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvironmentV2Request Request Object
type DeleteEnvironmentV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 环境的ID
	EnvId string `json:"env_id"`
}

func (o DeleteEnvironmentV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentV2Request", string(data)}, " ")
}
