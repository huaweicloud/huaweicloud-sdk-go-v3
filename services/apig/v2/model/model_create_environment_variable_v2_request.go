package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnvironmentVariableV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *EnvVariableReq `json:"body,omitempty"`
}

func (o CreateEnvironmentVariableV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentVariableV2Request struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentVariableV2Request", string(data)}, " ")
}
