package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportHostToEnvironmentRequest Request Object
type ImportHostToEnvironmentRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`

	Body *ImportHostToEnvironmentRequestBody `json:"body,omitempty"`
}

func (o ImportHostToEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportHostToEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"ImportHostToEnvironmentRequest", string(data)}, " ")
}
