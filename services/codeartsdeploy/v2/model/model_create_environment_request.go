package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentRequest Request Object
type CreateEnvironmentRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	Body *EnvironmentRequestBody `json:"body,omitempty"`
}

func (o CreateEnvironmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequest struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequest", string(data)}, " ")
}
