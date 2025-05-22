package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSecretsRequest Request Object
type CheckSecretsRequest struct {
	Body *CheckSecretsRequestBody `json:"body,omitempty"`
}

func (o CheckSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSecretsRequest struct{}"
	}

	return strings.Join([]string{"CheckSecretsRequest", string(data)}, " ")
}
