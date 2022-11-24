package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnvironmentV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *EnvCreate `json:"body,omitempty"`
}

func (o CreateEnvironmentV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentV2Request", string(data)}, " ")
}
