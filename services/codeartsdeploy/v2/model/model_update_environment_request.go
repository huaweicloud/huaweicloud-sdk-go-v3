package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvironmentRequest Request Object
type UpdateEnvironmentRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`

	Body *EnvironmentRequest `json:"body,omitempty"`
}

func (o UpdateEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentRequest", string(data)}, " ")
}
