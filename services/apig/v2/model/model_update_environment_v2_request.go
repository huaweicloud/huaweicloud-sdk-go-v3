package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEnvironmentV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 环境的ID，可通过查询环境信息获取该ID

	EnvId string `json:"env_id"`

	Body *EnvReq `json:"body,omitempty"`
}

func (o UpdateEnvironmentV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentV2Request", string(data)}, " ")
}
