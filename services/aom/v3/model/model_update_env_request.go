package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvRequest Request Object
type UpdateEnvRequest struct {

	// 环境id
	EnvironmentId string `json:"environment_id"`

	Body *EnvParam `json:"body,omitempty"`
}

func (o UpdateEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnvRequest", string(data)}, " ")
}
