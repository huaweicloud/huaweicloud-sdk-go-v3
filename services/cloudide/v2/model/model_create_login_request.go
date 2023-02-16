package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLoginRequest struct {
	Body *LoginSchema `json:"body,omitempty"`
}

func (o CreateLoginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginRequest struct{}"
	}

	return strings.Join([]string{"CreateLoginRequest", string(data)}, " ")
}
