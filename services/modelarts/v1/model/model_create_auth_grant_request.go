package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthGrantRequest Request Object
type CreateAuthGrantRequest struct {
	Body *AuthGrantRequest `json:"body,omitempty"`
}

func (o CreateAuthGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthGrantRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthGrantRequest", string(data)}, " ")
}
