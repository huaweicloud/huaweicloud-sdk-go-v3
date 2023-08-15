package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserResponse Response Object
type CreateUserResponse struct {
	User           *CreateUserResult `json:"user,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}
