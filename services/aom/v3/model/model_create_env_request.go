package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvRequest Request Object
type CreateEnvRequest struct {
	Body *EnvParam `json:"body,omitempty"`
}

func (o CreateEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvRequest struct{}"
	}

	return strings.Join([]string{"CreateEnvRequest", string(data)}, " ")
}
