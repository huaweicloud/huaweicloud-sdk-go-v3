package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizationResponse Response Object
type CreateAuthorizationResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthorizationResponse", string(data)}, " ")
}
