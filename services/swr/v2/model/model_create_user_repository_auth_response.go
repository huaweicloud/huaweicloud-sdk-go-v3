package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserRepositoryAuthResponse Response Object
type CreateUserRepositoryAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateUserRepositoryAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRepositoryAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateUserRepositoryAuthResponse", string(data)}, " ")
}
