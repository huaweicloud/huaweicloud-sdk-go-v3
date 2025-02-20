package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportSecretsRequest Request Object
type BatchImportSecretsRequest struct {
	Body *ImportSecretsRequest `json:"body,omitempty"`
}

func (o BatchImportSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportSecretsRequest struct{}"
	}

	return strings.Join([]string{"BatchImportSecretsRequest", string(data)}, " ")
}
