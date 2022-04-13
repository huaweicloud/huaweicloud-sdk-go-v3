package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKeyRequest struct {
	Body *CreateKeyRequestBody `json:"body,omitempty"`
}

func (o CreateKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyRequest", string(data)}, " ")
}
