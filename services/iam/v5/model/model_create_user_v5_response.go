package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserV5Response Response Object
type CreateUserV5Response struct {
	User           *User `json:"user,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateUserV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserV5Response struct{}"
	}

	return strings.Join([]string{"CreateUserV5Response", string(data)}, " ")
}
