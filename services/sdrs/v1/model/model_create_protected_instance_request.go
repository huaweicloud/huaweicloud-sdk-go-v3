package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProtectedInstanceRequest struct {
	Body *CreateProtectedInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateProtectedInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceRequest", string(data)}, " ")
}
