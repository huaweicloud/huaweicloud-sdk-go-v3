package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentV2Request Request Object
type CreateEnvironmentV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
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
