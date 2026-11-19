package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrustedKeySpaceRequest Request Object
type CreateTrustedKeySpaceRequest struct {
	Body *CreateTrustedKeySpaceRequestBody `json:"body,omitempty"`
}

func (o CreateTrustedKeySpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrustedKeySpaceRequest struct{}"
	}

	return strings.Join([]string{"CreateTrustedKeySpaceRequest", string(data)}, " ")
}
